package MMU

import (
	"emu6502/ComputeUnit/PrivRAM"
	"emu6502/Logger"
	"emu6502/RAM"
	"emu6502/ROM"
)

const (
	RAM_id = iota
	PrivRAM_id
	ROM_id
)

type Mapping struct {
	virtStart uint16
	physStart uint32
	size      uint16

	backingStore int
}

func NewMapping(virtStart uint16, physStart uint32, size uint16, backingStore int) *Mapping {
	return &Mapping{virtStart, physStart, size, backingStore}
}

func DefaultMappings() []*Mapping {
	return []*Mapping{
		NewMapping(0x0000, 0x0000, 0x2000, PrivRAM_id),
		NewMapping(0x2000, 0x0000, 0x2000, RAM_id),
		NewMapping(0x4000, 0x2000, 0x4000, RAM_id),
		NewMapping(0x4020, 0x0000, 0xBFDF, ROM_id),
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

	if mappings[0].backingStore != PrivRAM_id {
		Logger.Fatalf("PrivRAM must be the first mapping")
	}

	// check that privram is only mapped once
	for i := 1; i < len(mappings); i++ {
		if mappings[i].backingStore == PrivRAM_id {
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
			case PrivRAM_id:
				result := m.privRAM.Read(address)
				Logger.Debugf("Read PrivRAM[%04X] = %02X", address, result)
				return result
			case RAM_id:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.ramAddressBus <- RAM.AddressBus{Rw: 'R', Data: physicalAddress}
				result := (<-*m.ramDataBus).Data
				Logger.Debugf("Read RAM[%04X] = %02X", physicalAddress, result)
				return result
			case ROM_id:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.romAddressBus <- ROM.AddressBus{Rw: 'R', Data: physicalAddress}
				result := (<-*m.romDataBus).Data
				Logger.Debugf("Read ROM[%04X] = %02X", physicalAddress, result)
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
	return uint16(m.GetByteAt(address))<<8 | uint16(m.GetByteAt(address+1))
}

// SetByteAt writes the given byte to the given address
func (m *MMU) SetByteAt(address uint16, data uint8) {
	for _, mapping := range m.mappings {
		if address >= mapping.virtStart && address < mapping.virtStart+mapping.size {
			switch mapping.backingStore {
			case PrivRAM_id:
				m.privRAM.Write(address, data)
			case RAM_id:
				*m.ramAddressBus <- RAM.AddressBus{Rw: 'W', Data: m.convertVirtualAddressIntoPhysicalAddress(address)}
				*m.ramDataBus <- RAM.DataBus{Data: data}
			case ROM_id:
				Logger.Errorf("Cannot write to ROM")
			}
		}
	}

	Logger.Errorf("Write to unmapped memory: 0x%04X, %d", address, data)
}

// SetWordAt writes the given word to the given address
// It uses the SetByteAt method to write the bytes at the given address
func (m *MMU) SetWordAt(address uint16, data uint16) {
	m.SetByteAt(address, uint8(data>>8))
	m.SetByteAt(address+1, uint8(data))
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
