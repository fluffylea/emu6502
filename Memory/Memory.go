package Memory

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

type Memory struct {
	// Actual Memory
	ram [ramSize]uint8
	rom [romSize]uint8

	addressBus *chan AddressBus
	dataBus    *chan DataBus
}

// NewMemory is the constructor for a new Memory
func NewMemory(addressBus *chan AddressBus, dataBus *chan DataBus) *Memory {
	return &Memory{
		ram:        [ramSize]uint8{},
		rom:        [romSize]uint8{},
		addressBus: addressBus,
		dataBus:    dataBus}
}

// Reset initializes the Memory to 0
/*func (m *Memory) Reset() {
	_, err := os.Stat("mem.bin")
	if os.IsNotExist(err) {
		var i uint16
		for i = 0; i < memSize; i++ {
			switch i {
			case 0x0600:
				m.data[i] = 0xCC
			case 0x0601:
				m.data[i] = 0xFA
			case 0xFFFC:
				m.data[i] = 0x00
			case 0xFFFD:
				m.data[i] = 0x06
			default:
				m.data[i] = 0
			}
		}
		file, err := os.Create("mem.bin")
		if err != nil {
			panic(err.Error())
		}
		_, err = file.Write(m.data[:])
		if err != nil {
			panic(err.Error())
		}
		err = file.Close()
		if err != nil {
			panic(err.Error())
		}
	} else {
		bytes, err := os.ReadFile("mem.bin")
		if err != nil {
			panic(err.Error())
		}
		copy(m.data[:], bytes)
	}
}*/

func (m *Memory) Reset() {
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
func (m *Memory) Run() {
	for command := range *m.addressBus {
		if command.Rw == 'W' {
			m.handleMemoryWrite(command.Data, (<-*m.dataBus).Data)
		} else if command.Rw == 'R' {
			*m.dataBus <- DataBus{Data: m.handleMemoryRead(command.Data)}
		}
	}
}

func (m *Memory) handleMemoryWrite(location uint16, data uint8) {
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

func (m *Memory) handleMemoryRead(location uint16) uint8 {
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

func (m *Memory) writeToConsole(char uint8) {
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
