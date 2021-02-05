package CPU

// AddWithCarry adds two numbers with carry
// TODO AddWithCarry should probably return the entire processor status
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
	// If the result is bigger than one byte, set the carry flag
	c.ps.carry = addResult > 0xFF
	// Turn the result into a byte again
	result = uint8(addResult)

	return result
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
