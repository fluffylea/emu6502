package main

import (
	"emu6502/ComputeUnit/CPU"
	"emu6502/ComputeUnit/MMU"
	"emu6502/GPU"
	"emu6502/Logger"
	"emu6502/RAM"
	"emu6502/ROM"
	"flag"
	"strings"
)

var romFilename string

func init() {
	loglevel := flag.String("loglevel", "debug", "Debug mode")
	romFilenamePtr := flag.String("rom", "hello.rom", "Path to the ROM file")

	flag.Parse()

	switch strings.ToLower(*loglevel) {
	case "debug":
		Logger.ActiveLogLevel = Logger.LogLevelDebug
	case "info":
		Logger.ActiveLogLevel = Logger.LogLevelInfo
	case "error":
		Logger.ActiveLogLevel = Logger.LogLevelError
	}

	romFilename = *romFilenamePtr
}

func main() {
	ram := RAM.NewRAM()
	rom := ROM.NewROM()
	gpu := GPU.NewGPU()

	mappings := MMU.DefaultMappings()
	mmu := MMU.NewMMU(mappings, ram, rom, gpu)
	cpu := CPU.NewCPU(mmu)

	gpu.Reset()
	ram.Reset()
	rom.Reset(romFilename)

	go gpu.Run()
	go ram.Run()
	go rom.Run()

	cpu.Reset()
	cpu.Run()
}
