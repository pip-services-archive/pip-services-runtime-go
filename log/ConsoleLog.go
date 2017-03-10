package log

import (
	"github.com/pip-services/pip-services-runtime-go"
	"io"
	"os"
	"time"
)

type ConsoleLog struct {
	AbstractLog
}

func NewConsoleLog(config *runtime.DynamicMap) *ConsoleLog {
	return &ConsoleLog{AbstractLog: *NewAbstractLog("ConsoleLog", config)}
}

func FormatConsoleMessage(level string, message []interface{}) string {
	now := time.Now().UTC().Format(time.RFC3339)
	return now + " " + level + " " + formatMessage(message) + "\n"
}

func (c *ConsoleLog) Fatal(message ...interface{}) {
	if c.level >= runtime.LogLevelFatal {
		io.WriteString(os.Stderr, FormatConsoleMessage("FATAL", message))
	}
}

func (c *ConsoleLog) Error(message ...interface{}) {
	if c.level >= runtime.LogLevelError {
		io.WriteString(os.Stderr, FormatConsoleMessage("ERROR", message))
	}
}

func (c *ConsoleLog) Warn(message ...interface{}) {
	if c.level >= runtime.LogLevelWarn {
		io.WriteString(os.Stdout, FormatConsoleMessage("WARN", message))
	}
}

func (c *ConsoleLog) Info(message ...interface{}) {
	if c.level >= runtime.LogLevelInfo {
		io.WriteString(os.Stdout, FormatConsoleMessage("INFO", message))
	}
}

func (c *ConsoleLog) Debug(message ...interface{}) {
	if c.level >= runtime.LogLevelDebug {
		io.WriteString(os.Stdout, FormatConsoleMessage("DEBUG", message))
	}
}

func (c *ConsoleLog) Trace(message ...interface{}) {
	if c.level >= runtime.LogLevelTrace {
		io.WriteString(os.Stdout, FormatConsoleMessage("TRACE", message))
	}
}
