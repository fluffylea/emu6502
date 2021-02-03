package CPU

// AddWithCarry adds two numbers with carry
// TODO AddWithCarry should probably return the entire processor status
func AddWithCarry(number1 uint8, number2 uint8, carry bool) (result uint8, newCarry bool) {
	// Convert to 16-Bit variables
	num1Word := uint16(number1)
	num2Word := uint16(number2)
	// If the carry flag is set, add one
	if carry {
		num1Word++
	}
	// Do the calculation
	addResult := num1Word + num2Word
	// If the result is bigger than one byte, set the carry flag
	if addResult > 0xFF {
		newCarry = true
	} else {
		newCarry = false
	}
	// Turn the result into a byte again
	result = uint8(addResult)

	return result, newCarry
}
