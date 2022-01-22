package main

import (
	"emu6502/BusUnit"
	"emu6502/ComputeUnit"
	"emu6502/Logger"
	"flag"
	"strings"
	"time"
)

var romFilename string
var runtimeLimit int64

func init() {
	loglevel := flag.String("loglevel", "debug", "Debug mode")
	romFilenamePtr := flag.String("rom", "hello.rom", "Path to the ROM `file`")
	runtimeLimitPtr := flag.Int64("runtime", 10000, "Limit the runtime to the given number of `seconds`")
	Logger.DebugListingFile = flag.String("listing", "", "Path to the listing `file`")
	Logger.DebugMappingFile = flag.String("mapping", "", "Path to the mapping `file`")

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
	runtimeLimit = *runtimeLimitPtr
}

func main() {
	busUnit := BusUnit.NewBusUnit()

	cu1 := ComputeUnit.NewComputeUnit(busUnit)
	time.Now()

	busUnit.Reset(romFilename)
	busUnit.Run()

	cu1.Reset()
	cu1.Run()

	time.Sleep(time.Second * time.Duration(runtimeLimit))
	Logger.Infof("System exceeded runtime limit")

	Logger.Infof("Shutting down")

	cu1.Halt()
	busUnit.Halt()

	Logger.Infof("Shutdown complete")
}
