package AddressMode

type AddressMode struct {
	SelectedMode string
}

func Accumulator() AddressMode {
	return AddressMode{"A"}
}

func IsAccumulator(mode AddressMode) bool {
	return mode.SelectedMode == "A"
}

func Absolut() AddressMode {
	return AddressMode{"abs"}
}

func IsAbsolut(mode AddressMode) bool {
	return mode.SelectedMode == "abs"
}

func AbsolutX() AddressMode {
	return AddressMode{"abs,X"}
}

func IsAbsolutX(mode AddressMode) bool {
	return mode.SelectedMode == "abs,X"
}

func AbsolutY() AddressMode {
	return AddressMode{"abs,Y"}
}

func IsAbsolutY(mode AddressMode) bool {
	return mode.SelectedMode == "abs,Y"
}

func Immediate() AddressMode {
	return AddressMode{"#"}
}

func IsImmediate(mode AddressMode) bool {
	return mode.SelectedMode == "#"
}

func Implied() AddressMode {
	return AddressMode{"impl"}
}

func IsImplied(mode AddressMode) bool {
	return mode.SelectedMode == "impl"
}

func Indirect() AddressMode {
	return AddressMode{"ind"}
}

func IsIndirect(mode AddressMode) bool {
	return mode.SelectedMode == "ind"
}

func IndirectX() AddressMode {
	return AddressMode{"ind,X"}
}

func IsIndirectX(mode AddressMode) bool {
	return mode.SelectedMode == "ind,X"
}

func IndirectY() AddressMode {
	return AddressMode{"ind,Y"}
}

func IsIndirectY(mode AddressMode) bool {
	return mode.SelectedMode == "ind,Y"
}

func Relative() AddressMode {
	return AddressMode{"rel"}
}

func IsRelative(mode AddressMode) bool {
	return mode.SelectedMode == "rel"
}

func ZeroPage() AddressMode {
	return AddressMode{"zpg"}
}

func IsZeroPage(mode AddressMode) bool {
	return mode.SelectedMode == "zpg"
}

func ZeroPageX() AddressMode {
	return AddressMode{"zpg,X"}
}

func IsZeroPageX(mode AddressMode) bool {
	return mode.SelectedMode == "zpg,X"
}

func ZeroPageY() AddressMode {
	return AddressMode{"zpg,Y"}
}

func IsZeroPageY(mode AddressMode) bool {
	return mode.SelectedMode == "zpg,Y"
}
