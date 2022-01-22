package BusUnit

import (
	"emu6502/Logger"
	"fmt"
	"sync"
)

type GPU struct {
	AddressBus chan AddressBus
	DataBus    chan DataBus
	halt       *sync.WaitGroup
}

func NewGPU(wg *sync.WaitGroup) *GPU {
	wg.Add(1)
	return &GPU{
		AddressBus: make(chan AddressBus),
		DataBus:    make(chan DataBus),
		halt:       wg,
	}
}

func (g *GPU) Reset() {
	Logger.Infof("GPU Reset")
}

func (g *GPU) Run() {
	Logger.Infof("GPU Run")
	for command := range g.AddressBus {
		if command.Rw == 'W' {
			g.handleMemoryWrite(command.Data, (<-g.DataBus).Data)
		} else if command.Rw == 'R' {
			g.DataBus <- DataBus{Data: g.handleMemoryRead(command.Data)}
		}
	}
	g.halt.Done()
}

func (g *GPU) Halt() {
	Logger.Infof("GPU Halt")
	close(g.AddressBus)
	close(g.DataBus)
}

func (g *GPU) handleMemoryWrite(location uint32, data uint8) {
	switch location {
	case 0x00:
		fmt.Printf("%c", data)
	default:
		Logger.Warnf("GPU Memory Write: %x %x", location, data)
	}
}

func (g *GPU) handleMemoryRead(location uint32) uint8 {
	Logger.Warnf("GPU Memory Read: %d", location)
	return 0
}
