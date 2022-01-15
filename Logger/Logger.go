package Logger

import (
	"log"

	"github.com/fatih/color"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

var ActiveLogLevel = LogLevelInfo

func Debugf(format string, args ...interface{}) {
	if ActiveLogLevel <= LogLevelDebug {
		color.Set(color.FgHiGreen)
		log.Printf("[DEBUG] "+format+"\n", args...)
		color.Unset()
	}
}

func Infof(format string, args ...interface{}) {
	if ActiveLogLevel <= LogLevelInfo {
		color.Set(color.FgHiYellow)
		log.Printf("[INFO] "+format+"\n", args...)
		color.Unset()
	}
}

func Warnf(format string, args ...interface{}) {
	if ActiveLogLevel <= LogLevelWarn {
		color.Set(color.FgHiMagenta)
		log.Printf("[WARN] "+format+"\n", args...)
		color.Unset()
	}
}

func Errorf(format string, args ...interface{}) {
	if ActiveLogLevel <= LogLevelError {
		color.Set(color.FgHiRed)
		log.Printf("[ERROR] "+format+"\n", args...)
		color.Unset()
	}
}

func Fatalf(format string, args ...interface{}) {
	color.Set(color.FgHiRed, color.Bold)
	defer color.Unset()
	log.Panicf("[FATAL] "+format+"\n", args...)
}
