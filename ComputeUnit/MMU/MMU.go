package MMU

import (
	"emu6502/ComputeUnit/PrivRAM"
	"emu6502/Logger"
	"emu6502/RAM"
	"emu6502/ROM"
	"fmt"
)

const (
	RamId = iota
	PrivramId
	RomId
)

type Mapping struct {
	virtStart uint16
	physStart uint32
	size      uint16

	backingStore int
}

func (m *Mapping) ToString() string {
	return fmt.Sprintf("Mapping: virtStart: 0x%04x, physStart: 0x%04x, size: 0x%04x, backingStore: %d", m.virtStart, m.physStart, m.size, m.backingStore)
}

func NewMapping(virtStart uint16, physStart uint32, size uint16, backingStore int) *Mapping {
	return &Mapping{virtStart, physStart, size, backingStore}
}

func DefaultMappings() []*Mapping {
	return []*Mapping{
		NewMapping(0x0000, 0x0000, 0x2000, PrivramId),
		NewMapping(0x2000, 0x0000, 0x2000, RamId),
		NewMapping(0x4000, 0x2000, 0x0020, RamId),
		NewMapping(0x4020, 0x0000, 0xBFDF, RomId),
	}
}

// verifyMapping checks that there are no overlay mappings
func verifyMapping(mappings []*Mapping) {
	for i := 0; i < len(mappings); i++ {
		for j := i + 1; j < len(mappings); j++ {
			if mappings[i].virtStart <= mappings[j].virtStart &&
				mappings[i].virtStart+mappings[i].size > mappings[j].virtStart {
				Logger.Fatalf("Overlapping mappings: \n%s\n%s", mappings[i].ToString(), mappings[j].ToString())
			}
		}
	}
}

type MMU struct {
	privRAM       *PrivRAM.PrivRAM
	ramAddressBus *chan RAM.AddressBus
	ramDataBus    *chan RAM.DataBus
	romAddressBus *chan ROM.AddressBus
	romDataBus    *chan ROM.DataBus

	mappings []*Mapping
}

func NewMMU(mappings []*Mapping, ram *RAM.RAM, rom *ROM.ROM) *MMU {
	if mappings == nil {
		mappings = DefaultMappings()
	}

	verifyMapping(mappings)

	if mappings[0].backingStore != PrivramId {
		Logger.Fatalf("PrivRAM must be the first mapping")
	}

	// check that privram is only mapped once
	for i := 1; i < len(mappings); i++ {
		if mappings[i].backingStore == PrivramId {
			Logger.Fatalf("PrivRAM must only be mapped at first block")
		}
	}

	return &MMU{
		privRAM:       PrivRAM.NewPrivRAM(mappings[0].size),
		ramAddressBus: &ram.AddressBus,
		ramDataBus:    &ram.DataBus,

		romAddressBus: &rom.AddressBus,
		romDataBus:    &rom.DataBus,

		mappings: mappings,
	}
}

// GetByteAt returns the byte that is in Memory at the given address
func (m *MMU) GetByteAt(address uint16) uint8 {
	for _, mapping := range m.mappings {
		if address >= mapping.virtStart && address < mapping.virtStart+mapping.size {
			switch mapping.backingStore {
			case PrivramId:
				result := m.privRAM.Read(address)
				Logger.Debugf("Read PrivRAM[%04x] = %02x", address, result)
				return result
			case RamId:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.ramAddressBus <- RAM.AddressBus{Rw: 'R', Data: physicalAddress}
				result := (<-*m.ramDataBus).Data
				Logger.Debugf("Read RAM[%04x:%04x] = %02x", address, physicalAddress, result)
				return result
			case RomId:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.romAddressBus <- ROM.AddressBus{Rw: 'R', Data: physicalAddress}
				result := (<-*m.romDataBus).Data
				Logger.Debugf("Read ROM[%04x:%04x] = %02x", address, physicalAddress, result)
				return result
			}
		}
	}

	Logger.Errorf("Read from unmapped memory: 0x%04X", address)
	return 0
}

// GetWordAt reads a word from the given address
// It uses the GetByteAt method to read the bytes at the given address
func (m *MMU) GetWordAt(address uint16) uint16 {
	return uint16(m.GetByteAt(address+1))<<8 | uint16(m.GetByteAt(address))
}

// SetByteAt writes the given byte to the given address
func (m *MMU) SetByteAt(address uint16, data uint8) {
	for _, mapping := range m.mappings {
		if address >= mapping.virtStart && address < mapping.virtStart+mapping.size {
			switch mapping.backingStore {
			case PrivramId:
				m.privRAM.Write(address, data)
			case RamId:
				*m.ramAddressBus <- RAM.AddressBus{Rw: 'W', Data: m.convertVirtualAddressIntoPhysicalAddress(address)}
				*m.ramDataBus <- RAM.DataBus{Data: data}
			case RomId:
				Logger.Errorf("Cannot write to ROM")
			}
		}
	}

	Logger.Errorf("Write to unmapped memory: 0x%04X, %d", address, data)
}

// SetWordAt writes the given word to the given address
// It uses the SetByteAt method to write the bytes at the given address
func (m *MMU) SetWordAt(address uint16, data uint16) {
	m.SetByteAt(address+1, uint8(data>>8))
	m.SetByteAt(address, uint8(data))
}

// convertVirtualAddressIntoPhysicalAddress converts a virtual address into a physical address
// virtual address should be mapped.
func (m *MMU) convertVirtualAddressIntoPhysicalAddress(address uint16) uint32 {
	for _, mapping := range m.mappings {
		if address >= mapping.virtStart && address < mapping.virtStart+mapping.size {
			return mapping.physStart + uint32(address-mapping.virtStart)
		}
	}

	Logger.Fatalf("Virtual address not mapped")
	return 0
}
