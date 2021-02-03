package CPU

import (
	"bufio"
	"emu6502/CPU/AddressMode"
	"emu6502/Memory"
	"fmt"
	"log"
	"os"
)

const RESET_VECTOR = 0xFFFC

type CPU struct {
	// Program Counter
	pc uint16
	// Stack Pointer
	sp uint8

	// Accumulator
	a uint8
	// Index Register X
	x uint8
	// Index Register Y
	y uint8

	ps struct {
		carry      bool // Carry Flag
		zero       bool // Zero Flag
		intDisable bool // Interrupt Disable
		decimal    bool // Decimal Mode
		brk        bool // Break Command
		overflow   bool // Overflow Flag
		negative   bool // Negative Flag
	}

	addressBus *chan Memory.AddressBus
	dataBus    *chan Memory.DataBus
}

// NewCPU is the constructor for a new CPU
func NewCPU(addressBus *chan Memory.AddressBus, dataBus *chan Memory.DataBus) *CPU {
	return &CPU{addressBus: addressBus, dataBus: dataBus}
}

// Reset resets the CPU and gets it ready for execution
func (c *CPU) Reset() {
	// Set PC to 0xFFFB so the JMP instruction reads from 0xFFFC and 0xFFFD
	c.pc = 0xFFFB
	c.JMP(AddressMode.Absolut())
}

// Run starts the execution of the CPU
func (c *CPU) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		opcode := c.GetByteAt(c.pc)
		switch opcode {
		case 0x69:
			c.ADC(AddressMode.Immediate())
		case 0x65:
			c.ADC(AddressMode.ZeroPage())
		case 0x75:
			c.ADC(AddressMode.ZeroPageX())
		case 0x6D:
			c.ADC(AddressMode.Absolut())
		case 0x7D:
			c.ADC(AddressMode.AbsolutX())
		case 0x79:
			c.ADC(AddressMode.AbsolutY())
		case 0x61:
			c.ADC(AddressMode.IndirectX())
		case 0x71:
			c.ADC(AddressMode.IndirectY())
		case 0x29:
			c.AND(AddressMode.Immediate())
		case 0x25:
			c.AND(AddressMode.ZeroPage())
		case 0x35:
			c.AND(AddressMode.ZeroPageX())
		case 0x2D:
			c.AND(AddressMode.Absolut())
		case 0x3D:
			c.AND(AddressMode.AbsolutX())
		case 0x39:
			c.AND(AddressMode.AbsolutY())
		case 0x21:
			c.AND(AddressMode.IndirectX())
		case 0x31:
			c.AND(AddressMode.IndirectY())
		case 0x0A:
			c.ASL(AddressMode.Accumulator())
		case 0x06:
			c.ASL(AddressMode.ZeroPage())
		case 0x16:
			c.ASL(AddressMode.ZeroPage())
		case 0x0E:
			c.ASL(AddressMode.Absolut())
		case 0x1E:
			c.ASL(AddressMode.AbsolutX())
		case 0x90:
			c.BCC(AddressMode.Relative())
		case 0xB0:
			c.BCS(AddressMode.Relative())
		case 0xF0:
			c.BEQ(AddressMode.Relative())
		case 0x24:
			c.BIT(AddressMode.ZeroPage())
		case 0x2C:
			c.BIT(AddressMode.Absolut())
		case 0x30:
			c.BMI(AddressMode.Relative())
		case 0xD0:
			c.BNE(AddressMode.Relative())
		case 0x10:
			c.BPL(AddressMode.Relative())
		case 0x00:
			c.BRK(AddressMode.Implied())
		case 0x50:
			c.BVC(AddressMode.Relative())
		case 0x70:
			c.BVS(AddressMode.Relative())
		case 0x18:
			c.CLC(AddressMode.Implied())
		case 0xD8:
			c.CLD(AddressMode.Implied())
		case 0x58:
			c.CLI(AddressMode.Implied())
		case 0xB8:
			c.CLV(AddressMode.Implied())
		case 0xC9:
			c.CMP(AddressMode.Immediate())
		case 0xC5:
			c.CMP(AddressMode.ZeroPage())
		case 0xD5:
			c.CMP(AddressMode.ZeroPageX())
		case 0xCD:
			c.CMP(AddressMode.Absolut())
		case 0xDD:
			c.CMP(AddressMode.AbsolutX())
		case 0xD9:
			c.CMP(AddressMode.AbsolutY())
		case 0xC1:
			c.CMP(AddressMode.IndirectX())
		case 0xD1:
			c.CMP(AddressMode.IndirectY())
		case 0xE0:
			c.CPX(AddressMode.Immediate())
		case 0xE4:
			c.CPX(AddressMode.ZeroPage())
		case 0xEC:
			c.CPX(AddressMode.Absolut())
		case 0xC0:
			c.CPY(AddressMode.Immediate())
		case 0xC4:
			c.CPY(AddressMode.ZeroPage())
		case 0xCC:
			c.CPY(AddressMode.Absolut())
		case 0xC6:
			c.DEC(AddressMode.ZeroPage())
		case 0xD6:
			c.DEC(AddressMode.ZeroPageX())
		case 0xCE:
			c.DEC(AddressMode.Absolut())
		case 0xDE:
			c.DEC(AddressMode.AbsolutX())
		case 0xCA:
			c.DEX(AddressMode.Implied())
		case 0x88:
			c.DEY(AddressMode.Implied())
		case 0x49:
			c.EOR(AddressMode.Immediate())
		case 0x45:
			c.EOR(AddressMode.ZeroPage())
		case 0x55:
			c.EOR(AddressMode.ZeroPageX())
		case 0x4D:
			c.EOR(AddressMode.Absolut())
		case 0x5D:
			c.EOR(AddressMode.AbsolutX())
		case 0x59:
			c.EOR(AddressMode.AbsolutY())
		case 0x41:
			c.EOR(AddressMode.IndirectX())
		case 0x51:
			c.EOR(AddressMode.IndirectY())
		case 0xE6:
			c.INC(AddressMode.ZeroPage())
		case 0xF6:
			c.INC(AddressMode.ZeroPageX())
		case 0xEE:
			c.INC(AddressMode.Absolut())
		case 0xFE:
			c.INC(AddressMode.AbsolutX())
		case 0xE8:
			c.INX(AddressMode.Implied())
		case 0xC8:
			c.INY(AddressMode.Implied())
		case 0x4C:
			c.JMP(AddressMode.Absolut())
		case 0x6C:
			c.JMP(AddressMode.Indirect())
		case 0x20:
			c.JSR(AddressMode.Absolut())
		case 0xA9:
			c.LDA(AddressMode.Immediate())
		case 0xA5:
			c.LDA(AddressMode.ZeroPage())
		case 0xB5:
			c.LDA(AddressMode.ZeroPageX())
		case 0xAD:
			c.LDA(AddressMode.Absolut())
		case 0xBD:
			c.LDA(AddressMode.AbsolutX())
		case 0xB9:
			c.LDA(AddressMode.AbsolutY())
		case 0xA1:
			c.LDA(AddressMode.IndirectX())
		case 0xB1:
			c.LDA(AddressMode.IndirectY())
		case 0xA2:
			c.LDX(AddressMode.Immediate())
		case 0xA6:
			c.LDX(AddressMode.ZeroPage())
		case 0xB6:
			c.LDX(AddressMode.ZeroPageY())
		case 0xAE:
			c.LDX(AddressMode.Absolut())
		case 0xBE:
			c.LDX(AddressMode.AbsolutY())
		case 0xA0:
			c.LDY(AddressMode.Immediate())
		case 0xA4:
			c.LDY(AddressMode.ZeroPage())
		case 0xB4:
			c.LDY(AddressMode.ZeroPageX())
		case 0xAC:
			c.LDY(AddressMode.Absolut())
		case 0xBC:
			c.LDY(AddressMode.AbsolutX())
		case 0x4A:
			c.LSR(AddressMode.Accumulator())
		case 0x46:
			c.LSR(AddressMode.ZeroPage())
		case 0x56:
			c.LSR(AddressMode.ZeroPageX())
		case 0x4E:
			c.LSR(AddressMode.Absolut())
		case 0x5E:
			c.LSR(AddressMode.AbsolutX())
		case 0xEA:
			c.NOP(AddressMode.Implied())
		case 0x09:
			c.ORA(AddressMode.Immediate())
		case 0x05:
			c.ORA(AddressMode.ZeroPage())
		case 0x15:
			c.ORA(AddressMode.ZeroPageX())
		case 0x0D:
			c.ORA(AddressMode.Absolut())
		case 0x1D:
			c.ORA(AddressMode.AbsolutX())
		case 0x19:
			c.ORA(AddressMode.AbsolutY())
		case 0x01:
			c.ORA(AddressMode.IndirectX())
		case 0x11:
			c.ORA(AddressMode.IndirectY())
		case 0x48:
			c.PHA(AddressMode.Implied())
		case 0x08:
			c.PHP(AddressMode.Implied())
		case 0x68:
			c.PLA(AddressMode.Implied())
		case 0x28:
			c.PLP(AddressMode.Implied())
		case 0x2A:
			c.ROL(AddressMode.Accumulator())
		case 0x26:
			c.ROL(AddressMode.ZeroPage())
		case 0x36:
			c.ROL(AddressMode.ZeroPageX())
		case 0x2E:
			c.ROL(AddressMode.Absolut())
		case 0x3E:
			c.ROL(AddressMode.AbsolutX())
		case 0x6A:
			c.ROR(AddressMode.Accumulator())
		case 0x66:
			c.ROR(AddressMode.ZeroPage())
		case 0x76:
			c.ROR(AddressMode.ZeroPageX())
		case 0x6E:
			c.ROR(AddressMode.Absolut())
		case 0x7E:
			c.ROR(AddressMode.AbsolutX())
		case 0x40:
			c.RTI(AddressMode.Implied())
		case 0x60:
			c.RTS(AddressMode.Implied())
		case 0xE9:
			c.SBC(AddressMode.Immediate())
		case 0xE5:
			c.SBC(AddressMode.ZeroPage())
		case 0xF5:
			c.SBC(AddressMode.ZeroPageX())
		case 0xED:
			c.SBC(AddressMode.Absolut())
		case 0xFD:
			c.SBC(AddressMode.AbsolutX())
		case 0xF9:
			c.SBC(AddressMode.AbsolutY())
		case 0xE1:
			c.SBC(AddressMode.IndirectX())
		case 0xF1:
			c.SBC(AddressMode.IndirectY())
		case 0x38:
			c.SEC(AddressMode.Implied())
		case 0xF8:
			c.SED(AddressMode.Implied())
		case 0x78:
			c.SEI(AddressMode.Implied())
		case 0x85:
			c.STA(AddressMode.ZeroPage())
		case 0x95:
			c.STA(AddressMode.ZeroPageX())
		case 0x8D:
			c.STA(AddressMode.Absolut())
		case 0x9D:
			c.STA(AddressMode.AbsolutX())
		case 0x99:
			c.STA(AddressMode.AbsolutY())
		case 0x81:
			c.STA(AddressMode.IndirectX())
		case 0x91:
			c.STA(AddressMode.IndirectY())
		case 0x86:
			c.STX(AddressMode.ZeroPage())
		case 0x96:
			c.STX(AddressMode.ZeroPageY())
		case 0x8E:
			c.STX(AddressMode.Absolut())
		case 0x84:
			c.STY(AddressMode.ZeroPage())
		case 0x94:
			c.STY(AddressMode.ZeroPageX())
		case 0x8C:
			c.STY(AddressMode.Absolut())
		case 0xAA:
			c.TAX(AddressMode.Implied())
		case 0xA8:
			c.TAY(AddressMode.Implied())
		case 0xBA:
			c.TSX(AddressMode.Implied())
		case 0x8A:
			c.TXA(AddressMode.Implied())
		case 0x9A:
			c.TXS(AddressMode.Implied())
		case 0x98:
			c.TYA(AddressMode.Implied())
		}
		log.Println(c.ToString())

		if !scanner.Scan() {
			break
		}
	}
}

func (c *CPU) ToString() string {
	return fmt.Sprintf("PC: 0x%04x; SP: 0x%02x; A: 0x%02x; X: 0x%02x; Y: 0x%02x", c.pc, c.sp, c.a, c.x, c.y)
}
