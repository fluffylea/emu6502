package ROM

import (
	logger "emu6502/Logger"
	"os"
)

const romSize uint32 = 0xBFE0

type ROM struct {
	// Actual Memory
	rom [romSize]uint8

	AddressBus chan AddressBus
	DataBus    chan DataBus
}

// NewROM is the constructor for a new Memory
func NewROM() *ROM {
	return &ROM{
		rom:        [romSize]uint8{},
		AddressBus: make(chan AddressBus),
		DataBus:    make(chan DataBus)}
}

func (m *ROM) Reset(filename string) {
	logger.Infof("Loading ROM from %s", filename)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		logger.Fatalf("Please add a '%s' file", filename)
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		logger.Fatalf("Cannot read '%s'", filename)
	}

	logger.Infof("ROM Reset")
	copy(m.rom[:], bytes)
}

// Run starts with execution of the Memory
func (m *ROM) Run() {
	logger.Infof("ROM Run")
	for command := range m.AddressBus {
		if command.Rw == 'W' {
			m.handleMemoryWrite(command.Data, (<-m.DataBus).Data)
		} else if command.Rw == 'R' {
			m.DataBus <- DataBus{Data: m.handleMemoryRead(command.Data)}
		}
	}
}

func (m *ROM) handleMemoryWrite(location uint32, data uint8) {
	m.rom[location] = data
}

func (m *ROM) handleMemoryRead(location uint32) uint8 {
	return m.rom[location]
}
