package RAM

type AddressBus struct {
	// The Read/Write Byte indicates if the CPU wants to read from or write to memory
	Rw byte
	// actual contents
	Data uint32
}

type DataBus struct {
	// actual contents
	Data uint8
}
