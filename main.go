package main

import (
	"emu6502/ComputeUnit/CPU"
	"emu6502/ComputeUnit/MMU"
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
		Logger.ActiveLogLevel = Logger.LOG_LEVEL_DEBUG
	case "info":
		Logger.ActiveLogLevel = Logger.LOG_LEVEL_INFO
	case "error":
		Logger.ActiveLogLevel = Logger.LOG_LEVEL_ERROR
	}

	romFilename = *romFilenamePtr
}

func main() {
	ram := RAM.NewRAM()
	rom := ROM.NewROM()

	mappings := MMU.DefaultMappings()
	mmu := MMU.NewMMU(mappings, ram, rom)
	cpu := CPU.NewCPU(mmu)

	ram.Reset()
	rom.Reset(romFilename)

	go ram.Run()
	go rom.Run()

	cpu.Reset()
	cpu.Run()
}
