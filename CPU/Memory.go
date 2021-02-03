package CPU

import (
	"emu6502/Memory"
)

// GetNextByte returns the byte that is next in Memory according
// to the PC.
// Does NOT modify the PC
func (c *CPU) GetNextByte() uint8 {
	data := c.GetByteAt(c.pc + 1)
	return data
}

// GetNextWord returns the word that is next in Memory according
// to the PC
// Does NOT modify the PC
func (c *CPU) GetNextWord() uint16 {
	data := c.GetWordAt(c.pc + 1)
	return data
}

// GetByteAt returns the byte that is in Memory at the given address
func (c *CPU) GetByteAt(address uint16) uint8 {
	*c.addressBus <- Memory.AddressBus{Rw: 'R', Data: address}
	return (<-*c.dataBus).Data
}

// GetWordAt returns the word that is in Memory at the given address
func (c *CPU) GetWordAt(address uint16) uint16 {
	*c.addressBus <- Memory.AddressBus{Rw: 'R', Data: address}
	dataLow := (<-*c.dataBus).Data
	*c.addressBus <- Memory.AddressBus{Rw: 'R', Data: address + 1}
	dataHigh := (<-*c.dataBus).Data
	return CombineLowHigh(dataLow, dataHigh)
}

// SetByteAt sets the byte that is in Memory at the given address
func (c *CPU) SetByteAt(address uint16, data uint8) {
	*c.addressBus <- Memory.AddressBus{Rw: 'W', Data: address}
	*c.dataBus <- Memory.DataBus{Data: data}
}

// SetWordAt sets the word that is in Memory at the given address
func (c *CPU) SetWordAt(address uint16, data uint16) {
	dataLow, dataHigh := SplitLowHigh(data)
	*c.addressBus <- Memory.AddressBus{Rw: 'W', Data: address}
	*c.dataBus <- Memory.DataBus{Data: dataLow}
	*c.addressBus <- Memory.AddressBus{Rw: 'W', Data: address + 1}
	*c.dataBus <- Memory.DataBus{Data: dataHigh}
}

// CombineLowHigh combines a Low and a High byte into one Word
func CombineLowHigh(low uint8, high uint8) (combined uint16) {
	combined = uint16(high)<<8 + uint16(low)
	return combined
}

// SplitLowHigh splits a Word into a Low and a High byte
func SplitLowHigh(address uint16) (low uint8, high uint8) {
	low = uint8(address)
	high = uint8(address >> 8)
	return
}
