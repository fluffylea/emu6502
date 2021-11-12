package main

import (
	"emu6502/CPU"
	"emu6502/Memory"
)

func main() {
	var addrBus chan Memory.AddressBus = make(chan Memory.AddressBus)
	var dataBus chan Memory.DataBus = make(chan Memory.DataBus)

	mem := Memory.NewMemory(&addrBus, &dataBus, true)
	mem.Reset()
	go mem.Run()

	cpu := CPU.NewCPU(&addrBus, &dataBus)
	cpu.Reset()
	cpu.Run()
}
