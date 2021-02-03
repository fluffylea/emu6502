package Memory

import (
	"os"
)

const memSize uint16 = 0xFFFF

type Memory struct {
	// Actual Memory
	data [memSize]uint8

	addressBus *chan AddressBus
	dataBus    *chan DataBus
}

// NewMemory is the constructor for a new Memory
func NewMemory(addressBus *chan AddressBus, dataBus *chan DataBus) *Memory {
	return &Memory{data: [memSize]uint8{},
		addressBus: addressBus,
		dataBus:    dataBus}
}

// Reset initializes the Memory to 0
func (m *Memory) Reset() {
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
		defer file.Close()
		_, err = file.Write(m.data[:])
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
}

// Run starts with execution of the Memory
func (m *Memory) Run() {
	for command := range *m.addressBus {
		if command.Rw == 'W' {
			m.data[command.Data] = (<-*m.dataBus).Data
		} else if command.Rw == 'R' {
			*m.dataBus <- DataBus{Data: m.data[command.Data]}
		}
	}
}
