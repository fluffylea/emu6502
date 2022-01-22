package ComputeUnit

import (
	"emu6502/BusUnit"
	"emu6502/ComputeUnit/CPU"
	"emu6502/ComputeUnit/MMU"
	"emu6502/Logger"
	"sync"
)

type ComputeUnit struct {
	cpu *CPU.CPU
	mmu *MMU.MMU

	wg *sync.WaitGroup
}

func NewComputeUnit(busUnit *BusUnit.BusUnit) *ComputeUnit {
	wg := sync.WaitGroup{}
	mappings := MMU.DefaultMappings()

	connections := make([]*BusUnit.Connection, 3)
	connections[MMU.RamId] = &BusUnit.Connection{AddressBus: &busUnit.RAM.AddressBus, DataBus: &busUnit.RAM.DataBus}
	connections[MMU.RomId] = &BusUnit.Connection{AddressBus: &busUnit.ROM.AddressBus, DataBus: &busUnit.ROM.DataBus}
	connections[MMU.GpuId] = &BusUnit.Connection{AddressBus: &busUnit.GPU.AddressBus, DataBus: &busUnit.GPU.DataBus}
	mmu := MMU.NewMMU(mappings, connections)
	cpu := CPU.NewCPU(mmu, &wg)
	wg.Add(1)

	return &ComputeUnit{
		cpu: cpu,
		mmu: mmu,
		wg:  &wg,
	}
}

func (cu *ComputeUnit) Reset() {
	Logger.Infof("ComputeUnit Reset")
	cu.cpu.Reset()
}

func (cu *ComputeUnit) Run() {
	go cu.cpu.Run()
}

func (cu *ComputeUnit) Halt() {
	cu.cpu.Halt()
	cu.wg.Wait()
}
