package RAM

import "emu6502/Logger"

const ramSize uint32 = 0xFFFF

type RAM struct {
	// Actual Memory
	ram [ramSize]uint8

	AddressBus chan AddressBus
	DataBus    chan DataBus
}

// NewRAM is the constructor for a new Memory
func NewRAM() *RAM {
	return &RAM{
		ram:        [ramSize]uint8{},
		AddressBus: make(chan AddressBus),
		DataBus:    make(chan DataBus)}
}

// Reset resets the Memory to its initial state
func (m *RAM) Reset() {
	Logger.Infof("RAM Reset")
	for i := range m.ram {
		m.ram[i] = 0
	}
}

// Run starts with execution of the Memory
func (m *RAM) Run() {
	Logger.Infof("RAM Run")

	for command := range m.AddressBus {
		if command.Rw == 'W' {
			m.handleMemoryWrite(command.Data, (<-m.DataBus).Data)
		} else if command.Rw == 'R' {
			m.DataBus <- DataBus{Data: m.handleMemoryRead(command.Data)}
		}
	}
}

func (m *RAM) handleMemoryWrite(location uint32, data uint8) {
	m.ram[location] = data
}

func (m *RAM) handleMemoryRead(location uint32) uint8 {
	return m.ram[location]
}
