package Logger

import (
	"fmt"
	"os"
	"time"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

const (
	ColorHiGreen   = "\u001b[32;1m"
	ColorHiYellow  = "\u001b[33;1m"
	ColorHiMagenta = "\u001b[35;1m"
	ColorHiRed     = "\u001b[31;1m"
	ColorHiRedBold = "\u001B[31;1m\u001b[1m"
	ColorReset     = "\u001b[0m"
)

var ActiveLogLevel = LogLevelInfo

func Debugf(format string, args ...interface{}) {
	if ActiveLogLevel > LogLevelDebug {
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(os.Stderr, "%s%s [DEBUG] %s%s\n", ColorHiGreen, t, fmt.Sprintf(format, args...), ColorReset)
}

func Infof(format string, args ...interface{}) {
	if ActiveLogLevel > LogLevelInfo {
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(os.Stderr, "%s%s [INFO] %s%s\n", ColorHiYellow, t, fmt.Sprintf(format, args...), ColorReset)
}

func Warnf(format string, args ...interface{}) {
	if ActiveLogLevel > LogLevelWarn {
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(os.Stderr, "%s%s [WARN] %s%s\n", ColorHiMagenta, t, fmt.Sprintf(format, args...), ColorReset)
}

func Errorf(format string, args ...interface{}) {
	if ActiveLogLevel > LogLevelError {
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(os.Stderr, "%s%s [ERROR] %s%s\n", ColorHiRed, t, fmt.Sprintf(format, args...), ColorReset)
}

func Fatalf(format string, args ...interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05")
	panic(fmt.Sprintf("%s%s [FATAL] %s%s\n", ColorHiRedBold, t, fmt.Sprintf(format, args...), ColorReset))
}
