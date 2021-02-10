package CPU

import (
	"emu6502/CPU/AddressMode"
	"log"
)

// ADC performs an add with carry
func (c *CPU) ADC(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// ADC #$nn
		c.a = c.AddWithCarry(c.a, c.GetNextByte())
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ADC $ll
		parameter := c.GetNextByte()
		c.a = c.AddWithCarry(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ADC $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.AddWithCarry(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ADC $hhll
		parameter := c.GetNextWord()
		c.a = c.AddWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ADC $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.AddWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// ADC $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.AddWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// ADC ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.AddWithCarry(c.a, c.GetByteAt(CombineLowHigh(c.x, highByte)))
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// ADC ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.AddWithCarry(c.a, c.GetByteAt(addr))
		c.pc += 2
	default:
		log.Printf("ERR: ADC %s is not valid\n", mode.SelectedMode)
	}
}

// AND performs an and with the accumulator
func (c *CPU) AND(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// AND #$nn
		c.a = c.GetNextByte() & c.a
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// AND $ll
		parameter := c.GetNextByte()
		c.a = c.GetByteAt(uint16(parameter)) & c.a
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// AND $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.GetByteAt(uint16(parameter)) & c.a
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// AND $hhll
		parameter := c.GetNextWord()
		c.a = c.GetByteAt(parameter) & c.a
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// AND $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.GetByteAt(parameter) & c.a
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// AND $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.GetByteAt(parameter) & c.a
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// AND ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.GetByteAt(CombineLowHigh(c.x, highByte)) & c.a
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// AND ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.GetByteAt(addr) & c.a
		c.pc += 2
	default:
		log.Printf("ERR: AND %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(c.a)
	c.CheckZeroAndSetFlag(c.a)
}

// ASL performs an arithmetic shift left
func (c *CPU) ASL(mode AddressMode.AddressMode) {
	var tmp uint8 = 0
	switch {
	case AddressMode.IsAccumulator(mode):
		// ASL
		c.a = c.ArithmeticShiftLeft(c.a)
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ASL $ll
		parameter := c.GetNextByte()
		tmp = c.ArithmeticShiftLeft(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ASL $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.ArithmeticShiftLeft(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ASL $hhll
		parameter := c.GetNextWord()
		tmp = c.ArithmeticShiftLeft(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ASL $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.ArithmeticShiftLeft(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: ASL %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(tmp)
	c.CheckZeroAndSetFlag(tmp)
}

// BCC branches on carry clear
// TODO: Implement BCC
func (c *CPU) BCC(mode AddressMode.AddressMode) {
	log.Printf("ERR: BBC %s is not implemented\n", mode.SelectedMode)
}

// BCS branches on carry set
// TODO: Implement BCS
func (c *CPU) BCS(mode AddressMode.AddressMode) {
	log.Printf("ERR: BCS %s is not implemented\n", mode.SelectedMode)
}

// BEQ branches on equal (zero flag set)
// TODO: Implement BEQ
func (c *CPU) BEQ(mode AddressMode.AddressMode) {
	log.Printf("ERR: BEQ %s is not implemented\n", mode.SelectedMode)
}

// BIT bit test
// TODO: Implement BIT
func (c *CPU) BIT(mode AddressMode.AddressMode) {
	log.Printf("ERR: BIT %s is not implemented\n", mode.SelectedMode)
}

// BMI branches on minus (negative flag set)
// TODO: Implement BMI
func (c *CPU) BMI(mode AddressMode.AddressMode) {
	log.Printf("ERR: BMI %s is not implemented\n", mode.SelectedMode)
}

// BNE branches on not equal (zero flag clear)
// TODO: Implement BNE
func (c *CPU) BNE(mode AddressMode.AddressMode) {
	log.Printf("ERR: BNE %s is not implemented\n", mode.SelectedMode)
}

// BPL branches on plus (negative flag clear)
// TODO: Implement BPL
func (c *CPU) BPL(mode AddressMode.AddressMode) {
	log.Printf("ERR: BPL %s is not implemented\n", mode.SelectedMode)
}

// BRK break / interrupt
// TODO: Implement BRK
func (c *CPU) BRK(mode AddressMode.AddressMode) {
	log.Printf("ERR: BRK %s is not implemented\n", mode.SelectedMode)
}

// BVC branches on overflow clear
// TODO: Implement BVC
func (c *CPU) BVC(mode AddressMode.AddressMode) {
	log.Printf("ERR: BVC %s is not implemented\n", mode.SelectedMode)
}

// BVS branches on overflow set
// TODO: Implement BVS
func (c *CPU) BVS(mode AddressMode.AddressMode) {
	log.Printf("ERR: BVS %s is not implemented\n", mode.SelectedMode)
}

// CLC clears the carry flag
func (c *CPU) CLC(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.carry = false
		c.pc++
	default:
		log.Printf("ERR: CLC %s is not valid\n", mode.SelectedMode)
	}
}

// CLD clears the decimal flag
func (c *CPU) CLD(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.decimal = false
		c.pc++
	default:
		log.Printf("ERR: CLD %s is not valid\n", mode.SelectedMode)
	}
}

// CLI clears interrupt disable
func (c *CPU) CLI(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.intDisable = false
		c.pc++
	default:
		log.Printf("ERR: CLI %s is not valid\n", mode.SelectedMode)
	}
}

// CLV clears the overflow flag
func (c *CPU) CLV(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.overflow = false
		c.pc++
	default:
		log.Printf("ERR: CLV %s is not valid\n", mode.SelectedMode)
	}
}

// CMP compares with the accumulator
func (c *CPU) CMP(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// CMP #$nn
		c.Compare(c.a, c.GetNextByte())
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// CMP $ll
		parameter := c.GetNextByte()
		c.Compare(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// CMP $ll, X
		parameter := c.GetNextByte() + c.x
		c.Compare(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// CMP $hhll
		parameter := c.GetNextWord()
		c.Compare(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// CMP $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.Compare(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// CMP $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.Compare(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// CMP ($ll,X)
		highByte := c.GetNextByte()
		c.Compare(c.a, c.GetByteAt(CombineLowHigh(c.x, highByte)))
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// CMP ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.Compare(c.a, c.GetByteAt(addr))
		c.pc += 2
	default:
		log.Printf("ERR: CMP %s is not valid\n", mode.SelectedMode)
	}
}

// CPX compares with X
func (c *CPU) CPX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// CPX #$nn
		c.Compare(c.x, c.GetNextByte())
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// CPX $ll
		parameter := c.GetNextByte()
		c.Compare(c.x, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// CPX $hhll
		parameter := c.GetNextWord()
		c.Compare(c.x, c.GetByteAt(parameter))
		c.pc += 3
	default:
		log.Printf("ERR: CPX %s is not valid\n", mode.SelectedMode)
	}
}

// CPY compares with Y
func (c *CPU) CPY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// CPY #$nn
		c.Compare(c.y, c.GetNextByte())
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// CPY $ll
		parameter := c.GetNextByte()
		c.Compare(c.y, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// CPY $hhll
		parameter := c.GetNextWord()
		c.Compare(c.y, c.GetByteAt(parameter))
		c.pc += 3
	default:
		log.Printf("ERR: CPY %s is not valid\n", mode.SelectedMode)
	}
}

// DEC decrements memory
func (c *CPU) DEC(mode AddressMode.AddressMode) {
	var tmp uint8
	switch {
	case AddressMode.IsZeroPage(mode):
		// DEC $ll
		parameter := c.GetNextByte()
		tmp = c.GetByteAt(uint16(parameter)) - 1
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// DEC $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.GetByteAt(uint16(parameter)) - 1
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// DEC $hhll
		parameter := c.GetNextWord()
		tmp = c.GetByteAt(parameter) - 1
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// DEC $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.GetByteAt(parameter) - 1
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: DEC %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(c.a)
	c.CheckZeroAndSetFlag(c.a)
}

// DEX decrements X
func (c *CPU) DEX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.x--
		c.pc++
	default:
		log.Printf("ERR: DEX %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.x)
	c.CheckNegativeAndSetFlag(c.x)
}

// DEY decrements Y
func (c *CPU) DEY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.y--
		c.pc++
	default:
		log.Printf("ERR: DEY %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.y)
	c.CheckNegativeAndSetFlag(c.y)
}

// EOR performs an exclusive or with the accumulator
func (c *CPU) EOR(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// EOR #$nn
		c.a = c.GetNextByte() ^ c.a
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// EOR $ll
		parameter := c.GetNextByte()
		c.a = c.GetByteAt(uint16(parameter)) ^ c.a
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// EOR $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.GetByteAt(uint16(parameter)) ^ c.a
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// EOR $hhll
		parameter := c.GetNextWord()
		c.a = c.GetByteAt(parameter) ^ c.a
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// EOR $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.GetByteAt(parameter) ^ c.a
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// EOR $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.GetByteAt(parameter) ^ c.a
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// EOR ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.GetByteAt(CombineLowHigh(c.x, highByte)) ^ c.a
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// EOR ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.GetByteAt(addr) ^ c.a
		c.pc += 2
	default:
		log.Printf("ERR: EOR %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(c.a)
	c.CheckZeroAndSetFlag(c.a)
}

// INC increments memory
func (c *CPU) INC(mode AddressMode.AddressMode) {
	var tmp uint8
	switch {
	case AddressMode.IsZeroPage(mode):
		// INC $ll
		parameter := c.GetNextByte()
		tmp = c.GetByteAt(uint16(parameter)) + 1
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// INC $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.GetByteAt(uint16(parameter)) + 1
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// INC $hhll
		parameter := c.GetNextWord()
		tmp = c.GetByteAt(parameter) + 1
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// INC $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.GetByteAt(parameter) + 1
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: INC %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(c.a)
	c.CheckZeroAndSetFlag(c.a)
}

// INX increments X
func (c *CPU) INX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.x++
		c.pc++
	default:
		log.Printf("ERR: INX %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.x)
	c.CheckNegativeAndSetFlag(c.x)
}

// INY increments Y
func (c *CPU) INY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.y++
		c.pc++
	default:
		log.Printf("ERR: INY %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.y)
	c.CheckNegativeAndSetFlag(c.y)
}

// JMP sets the program counter to the new value to continue
// processing somewhere else
func (c *CPU) JMP(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsAbsolut(mode):
		c.pc = c.GetNextWord()
	case AddressMode.IsIndirect(mode):
		addr := c.GetNextWord()
		c.pc = c.GetWordAt(addr)
	default:
		log.Printf("ERR: JMP %s is not valid\n", mode.SelectedMode)
	}
}

// JSR jumps to a subroutine
// TODO: Implement JSR
func (c *CPU) JSR(mode AddressMode.AddressMode) {
	log.Printf("ERR: JSR %s is not implemented\n", mode.SelectedMode)
}

// LDA loads a value into the accumulator
func (c *CPU) LDA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// LDA #$nn
		c.a = c.GetNextByte()
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// LDA $ll
		parameter := c.GetNextByte()
		c.a = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// LDA $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// LDA $hhll
		parameter := c.GetNextWord()
		c.a = c.GetByteAt(parameter)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// LDA $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.GetByteAt(parameter)
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// LDA $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.GetByteAt(parameter)
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// LDA ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.GetByteAt(CombineLowHigh(c.x, highByte))
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// LDA ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.GetByteAt(addr)
		c.pc += 2
	default:
		log.Printf("ERR: LDA %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.a)
	c.CheckNegativeAndSetFlag(c.a)
}

// LDX loads a value into X
func (c *CPU) LDX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// LDX #$nn
		c.x = c.GetNextByte()
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// LDX $ll
		parameter := c.GetNextByte()
		c.x = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsZeroPageY(mode):
		// LDX $ll, Y
		parameter := c.GetNextByte() + c.y
		c.x = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// LDX $hhll
		parameter := c.GetNextWord()
		c.x = c.GetByteAt(parameter)
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// LDX $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.x = c.GetByteAt(parameter)
		c.pc += 3
	default:
		log.Printf("ERR: LDX %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.x)
	c.CheckNegativeAndSetFlag(c.x)
}

// LDY loads a value into Y
func (c *CPU) LDY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// LDY #$nn
		c.y = c.GetNextByte()
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// LDY $ll
		parameter := c.GetNextByte()
		c.y = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// LDY $ll, X
		parameter := c.GetNextByte() + c.x
		c.y = c.GetByteAt(uint16(parameter))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// LDY $hhll
		parameter := c.GetNextWord()
		c.y = c.GetByteAt(parameter)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// LDY $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.y = c.GetByteAt(parameter)
		c.pc += 3
	default:
		log.Printf("ERR: LDY %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.y)
	c.CheckNegativeAndSetFlag(c.y)
}

// LSR performs a logical shift right
func (c *CPU) LSR(mode AddressMode.AddressMode) {
	var tmp uint8 = 0
	switch {
	case AddressMode.IsAccumulator(mode):
		// LSR
		c.a = c.LogicalShiftRight(c.a)
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// LSR $ll
		parameter := c.GetNextByte()
		tmp = c.LogicalShiftRight(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// LSR $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.LogicalShiftRight(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// LSR $hhll
		parameter := c.GetNextWord()
		tmp = c.LogicalShiftRight(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// LSR $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.LogicalShiftRight(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: LSR %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(tmp)
	c.CheckZeroAndSetFlag(tmp)
}

// NOP does nothing
func (c *CPU) NOP(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.pc++
	default:
		log.Printf("ERR: NOP %s is not valid\n", mode.SelectedMode)
	}
}

// ORA ors with the accumulator
func (c *CPU) ORA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// ORA #$nn
		c.a = c.GetNextByte() | c.a
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ORA $ll
		parameter := c.GetNextByte()
		c.a = c.GetByteAt(uint16(parameter)) | c.a
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ORA $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.GetByteAt(uint16(parameter)) | c.a
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ORA $hhll
		parameter := c.GetNextWord()
		c.a = c.GetByteAt(parameter) | c.a
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ORA $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.GetByteAt(parameter) | c.a
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// ORA $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.GetByteAt(parameter) | c.a
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// ORA ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.GetByteAt(CombineLowHigh(c.x, highByte)) | c.a
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// ORA ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.GetByteAt(addr) | c.a
		c.pc += 2
	default:
		log.Printf("ERR: ORA %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(c.a)
	c.CheckZeroAndSetFlag(c.a)
}

// PHA push accumulator to the stack
// TODO: Implement PHA
func (c *CPU) PHA(mode AddressMode.AddressMode) {
	log.Printf("ERR: PHA %s is not implemented\n", mode.SelectedMode)
}

// PHP pushes the processor status to the stack
// TODO: Implement PHP
func (c *CPU) PHP(mode AddressMode.AddressMode) {
	log.Printf("ERR: PHP %s is not implemented\n", mode.SelectedMode)
}

// PLA pulls accumulator from the stack
// TODO: Implement PLA
func (c *CPU) PLA(mode AddressMode.AddressMode) {
	log.Printf("ERR: PLA %s is not implemented\n", mode.SelectedMode)
}

// PLP pulls the processor status from the stack
// TODO: Implement PLP
func (c *CPU) PLP(mode AddressMode.AddressMode) {
	log.Printf("ERR: PLP %s is not implemented\n", mode.SelectedMode)
}

// ROL rotates left
func (c *CPU) ROL(mode AddressMode.AddressMode) {
	var tmp uint8 = 0
	switch {
	case AddressMode.IsAccumulator(mode):
		// ROL
		c.a = c.RotateLeft(c.a)
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ROL $ll
		parameter := c.GetNextByte()
		tmp = c.RotateLeft(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ROL $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.RotateLeft(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ROL $hhll
		parameter := c.GetNextWord()
		tmp = c.RotateLeft(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ROL $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.RotateLeft(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: ROL %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(tmp)
	c.CheckZeroAndSetFlag(tmp)
}

// ROR rotates right
func (c *CPU) ROR(mode AddressMode.AddressMode) {
	var tmp uint8 = 0
	switch {
	case AddressMode.IsAccumulator(mode):
		// ROR
		c.a = c.RotateRight(c.a)
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ROR $ll
		parameter := c.GetNextByte()
		tmp = c.RotateRight(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ROR $ll, X
		parameter := c.GetNextByte() + c.x
		tmp = c.RotateRight(c.GetByteAt(uint16(parameter)))
		c.SetByteAt(uint16(parameter), tmp)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ROR $hhll
		parameter := c.GetNextWord()
		tmp = c.RotateRight(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ROR $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		tmp = c.RotateRight(c.GetByteAt(parameter))
		c.SetByteAt(parameter, tmp)
		c.pc += 3
	default:
		log.Printf("ERR: ROR %s is not valid\n", mode.SelectedMode)
	}
	c.CheckNegativeAndSetFlag(tmp)
	c.CheckZeroAndSetFlag(tmp)
}

// RTI returns from interrupt
// TODO: Implement RTI
func (c *CPU) RTI(mode AddressMode.AddressMode) {
	log.Printf("ERR: RTI %s is not implemented\n", mode.SelectedMode)
}

// RTS returns from subroutine
// TODO: Implement RTS
func (c *CPU) RTS(mode AddressMode.AddressMode) {
	log.Printf("ERR: RTS %s is not implemented\n", mode.SelectedMode)
}

// SBC subtracts with carry
func (c *CPU) SBC(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// SBC #$nn
		c.a = c.SubtractWithCarry(c.a, c.GetNextByte())
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// SBC $ll
		parameter := c.GetNextByte()
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// SBC $ll, X
		parameter := c.GetNextByte() + c.x
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(uint16(parameter)))
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// SBC $hhll
		parameter := c.GetNextWord()
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// SBC $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// SBC $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(parameter))
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// SBC ($ll,X)
		highByte := c.GetNextByte()
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(CombineLowHigh(c.x, highByte)))
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// SBC ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a = c.SubtractWithCarry(c.a, c.GetByteAt(addr))
		c.pc += 2
	default:
		log.Printf("ERR: SBC %s is not valid\n", mode.SelectedMode)
	}
}

// SEC sets the carry flag
func (c *CPU) SEC(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.carry = true
		c.pc++
	default:
		log.Printf("ERR: SEC %s is not valid\n", mode.SelectedMode)
	}
}

// SED sets the decimal flag
func (c *CPU) SED(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.decimal = true
		c.pc++
		log.Println("Just remember... BCD Mode doesn't work")
	default:
		log.Printf("ERR: SED %s is not valid\n", mode.SelectedMode)
	}
}

// SEI sets the interrupt disable flag
func (c *CPU) SEI(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.intDisable = true
		c.pc++
	default:
		log.Printf("ERR: SEI %s is not valid\n", mode.SelectedMode)
	}
}

// STA stores the accumulator in memory
func (c *CPU) STA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsZeroPage(mode):
		// STA $ll
		addr := c.GetNextByte()
		c.SetByteAt(uint16(addr), c.a)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// STA $ll, X
		addr := c.GetNextByte() + c.x
		c.SetByteAt(uint16(addr), c.a)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// STA $hhll
		addr := c.GetNextWord()
		c.SetByteAt(addr, c.a)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// STA $hhll,X
		addr := c.GetNextWord() + uint16(c.x)
		c.SetByteAt(addr, c.a)
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// STA $hhll,Y
		addr := c.GetNextWord() + uint16(c.y)
		c.SetByteAt(addr, c.a)
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// STA ($ll,X)
		highByte := c.GetNextByte()
		c.SetByteAt(CombineLowHigh(c.x, highByte), c.a)
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// STA ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.SetByteAt(addr, c.a)
		c.pc += 2
	default:
		log.Printf("ERR: STA %s is not valid\n", mode.SelectedMode)
	}
}

// STX stores X in memory
func (c *CPU) STX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsZeroPage(mode):
		// STX $ll
		addr := c.GetNextByte()
		c.SetByteAt(uint16(addr), c.x)
		c.pc += 2
	case AddressMode.IsZeroPageY(mode):
		// STX $ll, Y
		addr := c.GetNextByte() + c.y
		c.SetByteAt(uint16(addr), c.x)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// STX $hhll
		addr := c.GetNextWord()
		c.SetByteAt(addr, c.x)
		c.pc += 3
	default:
		log.Printf("ERR: STX %s is not valid\n", mode.SelectedMode)
	}
}

// STY stores Y in memory
func (c *CPU) STY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsZeroPage(mode):
		// STY $ll
		addr := c.GetNextByte()
		c.SetByteAt(uint16(addr), c.y)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// STY $ll, X
		addr := c.GetNextByte() + c.x
		c.SetByteAt(uint16(addr), c.y)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// STY $hhll
		addr := c.GetNextWord()
		c.SetByteAt(addr, c.y)
		c.pc += 3
	default:
		log.Printf("ERR: STY %s is not valid\n", mode.SelectedMode)
	}
}

// TAX transfers the accumulator to X
func (c *CPU) TAX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.x = c.a
		c.pc++
	default:
		log.Printf("ERR: TAX %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.x)
	c.CheckNegativeAndSetFlag(c.x)
}

// TAY transfers the accumulator to Y
func (c *CPU) TAY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.y = c.a
		c.pc++
	default:
		log.Printf("ERR: TAY %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.y)
	c.CheckNegativeAndSetFlag(c.y)
}

// TSX transfers the stack pointer to X
func (c *CPU) TSX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.x = c.sp
		c.pc++
	default:
		log.Printf("ERR: TSX %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.x)
	c.CheckNegativeAndSetFlag(c.x)
}

// TXA transfers X to the accumulator
func (c *CPU) TXA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.a = c.x
		c.pc++
	default:
		log.Printf("ERR: TXA %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.a)
	c.CheckNegativeAndSetFlag(c.a)
}

// TXS transfers X to the stack pointer
func (c *CPU) TXS(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.sp = c.x
		c.pc++
	default:
		log.Printf("ERR: TXS %s is not valid\n", mode.SelectedMode)
	}
}

// TYA transfers Y to the accumulator
func (c *CPU) TYA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.a = c.y
		c.pc++
	default:
		log.Printf("ERR: TXA %s is not valid\n", mode.SelectedMode)
	}
	c.CheckZeroAndSetFlag(c.a)
	c.CheckNegativeAndSetFlag(c.a)
}
