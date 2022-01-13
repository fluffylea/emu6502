package PrivRAM

type PrivRAM struct {
	storage []byte
}

func NewPrivRAM(size uint16) *PrivRAM {
	return &PrivRAM{make([]byte, size)}
}

func (p *PrivRAM) Read(address uint16) byte {
	return p.storage[address]
}

func (p *PrivRAM) Write(address uint16, data byte) {
	p.storage[address] = data
}
