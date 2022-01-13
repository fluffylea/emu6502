package MMU

import (
	"emu6502/ComputeUnit/PrivRAM"
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
		NewMapping(0x2000, 0x2000, 0x2000, RAM_id),
		NewMapping(0x4000, 0x4000, 0x4000, RAM_id),
		NewMapping(0x4020, 0x0000, 0xBFE0, ROM_id),
	}
}

type MMU struct {
	privRAM *PrivRAM.PrivRAM
	ram     *RAM.RAM
	rom     *ROM.ROM

	mappings []*Mapping
}

func NewMMU(mappings []*Mapping, ram *RAM.RAM, rom *ROM.ROM) *MMU {
	if mappings == nil {
		mappings = DefaultMappings()
	}

	if mappings[0].backingStore != PrivRAM_id {
		panic("PrivRAM must be the first mapping")
	}

	// check that privram is only mapped once
	for i := 1; i < len(mappings); i++ {
		if mappings[i].backingStore == PrivRAM_id {
			panic("PrivRAM must only be mapped at first block")
		}
	}

	return &MMU{
		privRAM: PrivRAM.NewPrivRAM(mappings[0].size),
		ram:     RAM.NewRAM(),
		rom:     ROM.NewROM(),

		mappings: mappings,
	}
}
