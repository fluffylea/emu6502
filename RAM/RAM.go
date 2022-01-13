package RAM

import (
	"fmt"
	"os"
)

const ramSize uint16 = 0x2000
const romSize uint16 = 0xBFE0

const ramOffset uint16 = 0x0000
const vramOffset uint16 = 0x2000
const ioOffset uint16 = 0x4000
const romOffset uint16 = 0x4020

type RAM struct {
	// Actual Memory
	ram [ramSize]uint8
	rom [romSize]uint8

	addressBus chan AddressBus
	dataBus    chan DataBus
}

// NewRAM is the constructor for a new Memory
func NewRAM() *RAM {
	return &RAM{
		ram:        [ramSize]uint8{},
		rom:        [romSize]uint8{},
		addressBus: make(chan AddressBus),
		dataBus:    make(chan DataBus)}
}

func (m *RAM) Reset() {
	_, err := os.Stat("rom.bin")
	if os.IsNotExist(err) {
		panic("Please add a rom.bin file")
	}

	bytes, err := os.ReadFile("rom.bin")
	if err != nil {
		panic("Cannot read rom.bin")
	}

	copy(m.rom[:], bytes)
}

// Run starts with execution of the Memory
func (m *RAM) Run() {
	for command := range m.addressBus {
		if command.Rw == 'W' {
			m.handleMemoryWrite(command.Data, (<-m.dataBus).Data)
		} else if command.Rw == 'R' {
			m.dataBus <- DataBus{Data: m.handleMemoryRead(command.Data)}
		}
	}
}

func (m *RAM) handleMemoryWrite(location uint16, data uint8) {
	switch {
	case isRAM(location):
		m.ram[location] = data
	case isVRAM(location):
		m.writeToConsole(data)
	case isIO(location):
		//These writes are ignored
	case isROM(location):
		m.rom[location-romOffset] = data
	}
}

func (m *RAM) handleMemoryRead(location uint16) uint8 {
	switch {
	case isRAM(location):
		return m.ram[location]
	case isVRAM(location):
		return 0
	case isIO(location):
		return 0
	case isROM(location):
		return m.rom[location-romOffset]
	default:
		return 0
	}
}

func (m *RAM) writeToConsole(char uint8) {
	fmt.Printf("%c", char)
}

func isRAM(location uint16) bool {
	return location >= ramOffset && location < vramOffset
}

func isVRAM(location uint16) bool {
	return location >= vramOffset && location < ioOffset
}

func isIO(location uint16) bool {
	return location >= ioOffset && location < romOffset
}

func isROM(location uint16) bool {
	return location >= romOffset
}
