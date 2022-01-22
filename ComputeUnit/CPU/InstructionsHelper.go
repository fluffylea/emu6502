package CPU

import (
	"math/bits"
	"unsafe"
)

// Uint8ToInt8 does what they tell you not to do:
// Convert the data type without touching the bits
func Uint8ToInt8(i uint8) int8 {
	return *(*int8)(unsafe.Pointer(&i))
}

// SubtractWithCarry is a helper function for subtracting
func SubtractWithCarry(number1 uint8, number2 uint8, carry bool) (result uint8, newOverflow bool, newCarry bool) {
	// Convert the Carry flag into a usable number
	var c uint16 = 0
	if carry {
		c = 1
	}
	// Perform the calculation: Sum of the One's complement plus the carry flag
	var res = uint16(number1) + uint16(^number2) + c

	// Check if the calculation would land us outside the borders of an 8-Bit signed Int
	var resInt = int16(Uint8ToInt8(number1)) + int16(Uint8ToInt8(^number2)) + int16(c)
	newOverflow = resInt > 127 || resInt < -128

	// Check if the result doesn't fit in one Byte
	var c7 = uint8(res >> 8)
	newCarry = c7 > 0

	// Turn the result into a byte
	result = uint8(res)

	return
}

// RotateRight rotates 8 Bit and the carry bit right
func RotateRight(accu uint8, carry bool) (uint8, bool) {
	tmp := uint16(accu)
	if carry {
		tmp += 0x100
	}
	carry = tmp&0b00000001 == 0b00000001
	accu = uint8(bits.RotateLeft16(tmp, -1))
	return accu, carry
}

// RotateRight rotates 8 Bit and the carry bit right
func (c *CPU) RotateRight(value uint8) (result uint8) {
	result, c.ps.carry = RotateRight(value, c.ps.carry)
	return result
}

// RotateLeft rotates 8 bit and the carry bit left
func RotateLeft(accu uint8, carry bool) (uint8, bool) {
	tmp := uint16(accu) << 1
	if carry {
		tmp += 1
	}
	carry = tmp&0x100 == 0x100
	return uint8(tmp), carry
}

// RotateLeft rotates 8 bit and the carry bit left
func (c *CPU) RotateLeft(value uint8) (result uint8) {
	result, c.ps.carry = RotateLeft(value, c.ps.carry)
	return result
}

// LogicalShiftRight performs a shift right into the Carry
func LogicalShiftRight(accu uint8, carry bool) (uint8, bool) {
	tmp := uint16(accu) << 1
	if carry {
		tmp += 1
	}

	carry = tmp&0x100 == 0x100
	return uint8(tmp), carry
}

// LogicalShiftRight performs a shift right into the Carry
func (c *CPU) LogicalShiftRight(value uint8) (result uint8) {
	result, c.ps.carry = LogicalShiftRight(value, c.ps.carry)
	return result
}

// SubtractWithCarry subtracts two number with carry
func (c *CPU) SubtractWithCarry(number1 uint8, number2 uint8) (result uint8) {
	result, c.ps.overflow, c.ps.carry = SubtractWithCarry(number1, number2, c.ps.carry)

	c.CheckNegativeAndSetFlag(result)
	c.CheckZeroAndSetFlag(result)

	return result
}

// Compare compares two numbers and sets the Carry, Negative and Zero flags accordingly
// Separate function is necessary because CMP doesn't set the overflow flag
func (c *CPU) Compare(number1 uint8, number2 uint8) {
	var res uint8
	res, _, c.ps.carry = SubtractWithCarry(number1, number2, true)
	c.CheckNegativeAndSetFlag(res)
	c.CheckZeroAndSetFlag(res)
}

// AddWithCarry adds two numbers with carry
func (c *CPU) AddWithCarry(number1 uint8, number2 uint8) (result uint8) {
	// Convert to 16-Bit variables
	num1Word := uint16(number1)
	num2Word := uint16(number2)
	// If the carry flag is set, add one
	if c.ps.carry {
		num1Word++
	}
	// Do the calculation
	addResult := num1Word + num2Word

	// Set the overflow and carry flag
	// http://www.righto.com/2012/12/the-6502-overflow-flag-explained.html
	var c6 = ((number1 & 0b01111111) + (number2 & 0b01111111)) >> 7
	var c7 = uint8(addResult >> 8)
	c.ps.overflow = (c6^c7)&0b00000001 == 1
	c.ps.carry = c7 > 0
	// Turn the result into a byte again
	result = uint8(addResult)

	c.CheckZeroAndSetFlag(result)
	c.CheckNegativeAndSetFlag(result)

	return result
}

// ArithmeticShiftLeft performs an ASR and puts the shifted out bit into the carry flag
func (c *CPU) ArithmeticShiftLeft(number uint8) uint8 {
	c.ps.carry = number&0b10000000 > 0
	return number << 1
}

// CheckZeroAndSetFlag checks if the number is zero and sets the flag accordingly
func (c *CPU) CheckZeroAndSetFlag(number uint8) {
	c.ps.zero = number == 0
}

// CheckNegativeAndSetFlag mirrors the first bit of number into the negative flag
func (c *CPU) CheckNegativeAndSetFlag(number uint8) {
	c.ps.negative = ConvertUint8ToBits(number)[0]
}

// ConvertUint8ToBits converts a given uint8 number to a boolean array
func ConvertUint8ToBits(number uint8) (bits [8]bool) {
	for i := 7; i >= 0; i-- {
		bits[i] = number&0x1 == 0x1
		number = number >> 1
	}
	return bits
}

// ConvertBitsToUint8 converts a boolean array into an uint8 number
func ConvertBitsToUint8(bits [8]bool) (number uint8) {
	for i := 0; i < 8; i++ {
		number = number << 1
		if bits[i] {
			number += 1
		}
	}
	return number
}

// PushToStack pushes data onto the stack and decrements the stack pointer
func (c *CPU) PushToStack(data uint8) {
	c.SetByteAt(0x100+uint16(c.sp), data)
	c.sp--
}

// PullFromStack pulls data from the stack and increments the stack pointer
func (c *CPU) PullFromStack() (data uint8) {
	c.sp++
	data = c.GetByteAt(0x100 + uint16(c.sp))
	return data
}

// PushWordToStack pushes an entire word onto the stack
func (c *CPU) PushWordToStack(data uint16) {
	dataLow := uint8(data)
	dataHigh := uint8(data >> 8)
	c.PushToStack(dataHigh)
	c.PushToStack(dataLow)
}

// PullWordFromStack pulls an entire word from the stack
func (c *CPU) PullWordFromStack() (data uint16) {
	dataLow := uint16(c.PullFromStack())
	dataHigh := uint16(c.PullFromStack()) << 8

	data = dataLow + dataHigh
	return data
}
