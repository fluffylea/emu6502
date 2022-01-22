package BusUnit

import (
	"sync"
)

type BusUnit struct {
	ROM *ROM
	RAM *RAM
	GPU *GPU

	wg *sync.WaitGroup
}

func NewBusUnit() *BusUnit {
	wg := &sync.WaitGroup{}
	return &BusUnit{
		ROM: NewROM(wg),
		RAM: NewRAM(wg),
		GPU: NewGPU(wg),
		wg:  wg,
	}
}

func (bus *BusUnit) Reset(romFilename string) {
	bus.ROM.Reset(romFilename)
	bus.RAM.Reset()
	bus.GPU.Reset()
}

func (bus *BusUnit) Run() {
	go bus.ROM.Run()
	go bus.RAM.Run()
	go bus.GPU.Run()
}

func (bus *BusUnit) Halt() {
	bus.ROM.Halt()
	bus.RAM.Halt()
	bus.GPU.Halt()

	bus.wg.Wait()
}
