package CPU

import (
	"emu6502/CPU/AddressMode"
	"log"
)

// TODO: Fix flags everywhere

// ADC performs an add with carry
// TODO: When flags are fixed, fix this
func (c *CPU) ADC(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImmediate(mode):
		// ADC #$nn
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetNextByte(), c.ps.carry)
		c.pc += 2
	case AddressMode.IsZeroPage(mode):
		// ADC $ll
		parameter := c.GetNextByte()
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(uint16(parameter)), c.ps.carry)
		c.pc += 2
	case AddressMode.IsZeroPageX(mode):
		// ADC $ll, X
		parameter := c.GetNextByte() + c.x
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(uint16(parameter)), c.ps.carry)
		c.pc += 2
	case AddressMode.IsAbsolut(mode):
		// ADC $hhll
		parameter := c.GetNextWord()
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(parameter), c.ps.carry)
		c.pc += 3
	case AddressMode.IsAbsolutX(mode):
		// ADC $hhll,X
		parameter := c.GetNextWord() + uint16(c.x)
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(parameter), c.ps.carry)
		c.pc += 3
	case AddressMode.IsAbsolutY(mode):
		// ADC $hhll,Y
		parameter := c.GetNextWord() + uint16(c.y)
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(parameter), c.ps.carry)
		c.pc += 3
	case AddressMode.IsIndirectX(mode):
		// ADC ($ll,X)
		highByte := c.GetNextByte()
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(CombineLowHigh(c.x, highByte)), c.ps.carry)
		c.pc += 2
	case AddressMode.IsIndirectY(mode):
		// ADC ($ll),Y
		parameter := c.GetNextByte()
		addr := c.GetWordAt(uint16(parameter)) + uint16(c.y)
		c.a, c.ps.carry = AddWithCarry(c.a, c.GetByteAt(addr), c.ps.carry)
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
}

// ASL performs an arithmetic shift left
// TODO: Implement ASL
func (c *CPU) ASL(mode AddressMode.AddressMode) {
	log.Printf("ERR: ASL %s is not implemented\n", mode.SelectedMode)
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
// TODO: Implement CMP
func (c *CPU) CMP(mode AddressMode.AddressMode) {
	log.Printf("ERR: CMP %s is not implemented\n", mode.SelectedMode)
}

// CPX compares with X
// TODO: Implement CPX
func (c *CPU) CPX(mode AddressMode.AddressMode) {
	log.Printf("ERR: CPX %s is not implemented\n", mode.SelectedMode)
}

// CPY compares with Y
// TODO: Implement CPY
func (c *CPU) CPY(mode AddressMode.AddressMode) {
	log.Printf("ERR: CPY %s is not implemented\n", mode.SelectedMode)
}

// DEC decrements memory
// TODO: Implement DEC
func (c *CPU) DEC(mode AddressMode.AddressMode) {
	log.Printf("ERR: DEC %s is not implemented\n", mode.SelectedMode)
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
}

// EOR performs an exclusive or with the accumulator
// TODO: Implement EOR
func (c *CPU) EOR(mode AddressMode.AddressMode) {
	log.Printf("ERR: EOR %s is not implemented\n", mode.SelectedMode)
}

// INC increments memory
// TODO: Implement INC
func (c *CPU) INC(mode AddressMode.AddressMode) {
	log.Printf("ERR: INC %s is not implemented\n", mode.SelectedMode)
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
}

// LSR performs a logical shift right
// TODO: Implement LSR
func (c *CPU) LSR(mode AddressMode.AddressMode) {
	log.Printf("ERR: LSR %s is not implemented\n", mode.SelectedMode)
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
// TODO: Implement ORA
func (c *CPU) ORA(mode AddressMode.AddressMode) {
	log.Printf("ERR: ORA %s is not implemented\n", mode.SelectedMode)
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
// TODO: Implement ROL
func (c *CPU) ROL(mode AddressMode.AddressMode) {
	log.Printf("ERR: ROL %s is not implemented\n", mode.SelectedMode)
}

// ROR rotates right
// TODO: Implement ROR
func (c *CPU) ROR(mode AddressMode.AddressMode) {
	log.Printf("ERR: ROR %s is not implemented\n", mode.SelectedMode)
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
// TODO: Implement SBC
func (c *CPU) SBC(mode AddressMode.AddressMode) {
	log.Printf("ERR: SBC %s is not implemented\n", mode.SelectedMode)
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
		c.ps.zero = c.a == 0
		c.ps.negative = c.a&128 == 128
		c.x = c.a
		c.pc++
	default:
		log.Printf("ERR: TAX %s is not valid\n", mode.SelectedMode)
	}
}

// TAY transfers the accumulator to Y
func (c *CPU) TAY(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.zero = c.a == 0
		c.ps.negative = c.a&128 == 128
		c.y = c.a
		c.pc++
	default:
		log.Printf("ERR: TAY %s is not valid\n", mode.SelectedMode)
	}
}

// TSX transfers the stack pointer to X
func (c *CPU) TSX(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.zero = c.sp == 0
		c.ps.negative = c.sp&128 == 128
		c.x = c.sp
		c.pc++
	default:
		log.Printf("ERR: TSX %s is not valid\n", mode.SelectedMode)
	}
}

// TXA transfers X to the accumulator
func (c *CPU) TXA(mode AddressMode.AddressMode) {
	switch {
	case AddressMode.IsImplied(mode):
		c.ps.zero = c.x == 0
		c.ps.negative = c.x&128 == 128
		c.a = c.x
		c.pc++
	default:
		log.Printf("ERR: TXA %s is not valid\n", mode.SelectedMode)
	}
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
		c.ps.zero = c.y == 0
		c.ps.negative = c.y&128 == 128
		c.a = c.y
		c.pc++
	default:
		log.Printf("ERR: TXA %s is not valid\n", mode.SelectedMode)
	}
}