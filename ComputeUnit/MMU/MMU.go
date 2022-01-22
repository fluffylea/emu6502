package MMU

import (
	"emu6502/BusUnit"
	"emu6502/ComputeUnit/PrivRAM"
	"emu6502/Logger"
	"fmt"
)

const (
	RamId = iota
	RomId
	GpuId
	MmuId
	PrivramId
)

const mappingSize = 2 + 4 + 2 + 1

type Mapping struct {
	virtStart uint16
	physStart uint32
	size      uint16

	backingStore uint8
}

func getByteFromMapping(mappings []*Mapping, addr uint32) uint8 {
	// get number of mapping from addr and mapping size
	mappingNum := addr / mappingSize
	// get offset from addr
	offset := addr % mappingSize

	// get mapping
	mapping := mappings[mappingNum]

	switch offset {
	case 0:
		return uint8(mapping.virtStart)
	case 1:
		return uint8(mapping.virtStart >> 8)
	case 2:
		return uint8(mapping.physStart)
	case 3:
		return uint8(mapping.physStart >> 8)
	case 4:
		return uint8(mapping.physStart >> 16)
	case 5:
		return uint8(mapping.physStart >> 24)
	case 6:
		return uint8(mapping.size)
	case 7:
		return uint8(mapping.size >> 8)
	case 8:
		return mapping.backingStore
	default:
		Logger.Warnf("MMU: getByteFromMapping: invalid offset %d", offset)
		return 0
	}
}

func (m *Mapping) ToString() string {
	return fmt.Sprintf("Mapping: virtStart: 0x%04x, physStart: 0x%04x, size: 0x%04x, backingStore: %d", m.virtStart, m.physStart, m.size, m.backingStore)
}

func NewMapping(virtStart uint16, physStart uint32, size uint16, backingStore uint8) *Mapping {
	return &Mapping{virtStart, physStart, size, backingStore}
}

func DefaultMappings() []*Mapping {
	return []*Mapping{
		NewMapping(0x0000, 0x0000, 0x2000, PrivramId),
		NewMapping(0x2000, 0x0000, 0x1FE0, RamId),
		NewMapping(0x3FE0, 0x0000, 0x0020, MmuId),
		NewMapping(0x4000, 0x0000, 0x0020, GpuId),
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
	connections []*BusUnit.Connection
	privRAM     *PrivRAM.PrivRAM
	mappings    []*Mapping
}

func NewMMU(mappings []*Mapping, connections []*BusUnit.Connection) *MMU {
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
		connections: connections,
		privRAM:     PrivRAM.NewPrivRAM(mappings[0].size),
		mappings:    mappings,
	}
}

// GetByteAt returns the byte that is in Memory at the given address
func (m *MMU) GetByteAt(address uint16) uint8 {
	for _, mapping := range m.mappings {
		if address >= mapping.virtStart && address < mapping.virtStart+mapping.size {
			switch mapping.backingStore {
			case PrivramId:
				result := m.privRAM.Read(address)
				return result
			case RamId:
				fallthrough
			case RomId:
				fallthrough
			case GpuId:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.connections[mapping.backingStore].AddressBus <- BusUnit.AddressBus{Rw: 'R', Data: physicalAddress}
				result := (<-*m.connections[mapping.backingStore].DataBus).Data
				return result
			case MmuId:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				getByteFromMapping(m.mappings, physicalAddress)
			default:
				Logger.Fatalf("Unknown backing store: %d", mapping.backingStore)
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
				return
			case RamId:
				fallthrough
			case RomId:
				fallthrough
			case GpuId:
				physicalAddress := m.convertVirtualAddressIntoPhysicalAddress(address)
				*m.connections[mapping.backingStore].AddressBus <- BusUnit.AddressBus{Rw: 'W', Data: physicalAddress}
				*m.connections[mapping.backingStore].DataBus <- BusUnit.DataBus{Data: data}
				return
			case MmuId:
				Logger.Errorf("Writing to MMU not implemented")
				return
			default:
				Logger.Fatalf("Unknown backing store: %d", mapping.backingStore)
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
