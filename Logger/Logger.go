package Logger

import (
	"log"

	"github.com/fatih/color"
)

const (
	LOG_LEVEL_DEBUG = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

var ActiveLogLevel = LOG_LEVEL_INFO

func Debugf(format string, args ...interface{}) {
	if ActiveLogLevel <= LOG_LEVEL_DEBUG {
		color.Set(color.FgHiGreen)
		defer color.Unset()

		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func Infof(format string, args ...interface{}) {
	if ActiveLogLevel <= LOG_LEVEL_INFO {
		color.Set(color.FgHiYellow)
		defer color.Unset()

		log.Printf("[INFO] "+format+"\n", args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if ActiveLogLevel <= LOG_LEVEL_WARN {
		color.Set(color.FgHiMagenta)
		defer color.Unset()

		log.Printf("[WARN] "+format+"\n", args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if ActiveLogLevel <= LOG_LEVEL_ERROR {
		color.Set(color.FgHiRed)
		defer color.Unset()

		log.Printf("[ERROR] "+format+"\n", args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	color.Set(color.FgHiRed, color.Bold)
	defer color.Unset()

	log.Panicf("[FATAL] "+format+"\n", args...)
}