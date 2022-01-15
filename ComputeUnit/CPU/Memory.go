package CPU

// TODO: Those are just wrapper functions around mmu functions
//       The CPU Instructions should be refactored to directly access
//       the mmu.

// GetNextByte returns the byte that is next in Memory according
// to the PC.
// Does NOT modify the PC
func (c *CPU) GetNextByte() uint8 {
	return c.mmu.GetByteAt(c.pc + 1)
}

// GetNextWord returns the word that is next in Memory according
// to the PC
// Does NOT modify the PC
func (c *CPU) GetNextWord() uint16 {
	return c.mmu.GetWordAt(c.pc + 1)
}

// GetByteAt returns the byte that is in Memory at the given address
func (c *CPU) GetByteAt(address uint16) uint8 {
	return c.mmu.GetByteAt(address)
}

// GetWordAt returns the word that is in Memory at the given address
func (c *CPU) GetWordAt(address uint16) uint16 {
	return c.mmu.GetWordAt(address)
}

// SetByteAt sets the byte that is in Memory at the given address
func (c *CPU) SetByteAt(address uint16, data uint8) {
	c.mmu.SetByteAt(address, data)
}

// SetWordAt sets the word that is in Memory at the given address
func (c *CPU) SetWordAt(address uint16, data uint16) {
	c.mmu.SetWordAt(address, data)
}

// CombineLowHigh combines a Low and a High byte into one Word
func CombineLowHigh(low uint8, high uint8) (combined uint16) {
	combined = uint16(high)<<8 + uint16(low)
	return combined
}
