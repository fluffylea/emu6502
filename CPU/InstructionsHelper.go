package CPU

import "unsafe"

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
	var res uint16 = uint16(number1) + uint16(^number2) + c

	// Check if the calculation would land us outside the borders of a 8 Bit signed Int
	var resInt = int16(Uint8ToInt8(number1)) + int16(Uint8ToInt8(^number2)) + int16(c)
	newOverflow = resInt > 127 || resInt < -128

	// Check if the result doesn't fit in one Byte
	var c7 uint8 = uint8(res >> 8)
	newCarry = c7 > 0

	// Turn the result into a byte
	result = uint8(res)

	return
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
	var c6 uint8 = ((number1 & 0b01111111) + (number2 & 0b01111111)) >> 7
	var c7 uint8 = uint8(addResult >> 8)
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

// ConvertBitsToUint8 converts a boolean array into a uint8 number
func ConvertBitsToUint8(bits [8]bool) (number uint8) {
	for i := 0; i < 8; i++ {
		number = number << 1
		if bits[i] == true {
			number += 1
		}
	}
	return number
}
