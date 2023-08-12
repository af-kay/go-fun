package extendedlogger

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

var prefixes = map[LogLevel]string{
	LogLevelError:   "ERROR",
	LogLevelWarning: "WARN",
	LogLevelInfo:    "INFO",
	LogLevelDebug:   "DEBUG",
	LogLevelTrace:   "TRACE",
}

type ExtendedLogger struct {
	*log.Logger
	level LogLevel
}

func (l *ExtendedLogger) println(level LogLevel, msg string) {
	if l.level < level {
		return
	}

	l.Logger.Println(prefixes[level] + " " + msg)
}

func (l *ExtendedLogger) Errorln(msg string) {
	l.println(LogLevelError, msg)
}
func (l *ExtendedLogger) Warnln(msg string) {
	l.println(LogLevelWarning, msg)
}
func (l *ExtendedLogger) Infoln(msg string) {
	l.println(LogLevelInfo, msg)
}
func (l *ExtendedLogger) Debugln(msg string) {
	l.println(LogLevelDebug, msg)
}
func (l *ExtendedLogger) Traceln(msg string) {
	l.println(LogLevelTrace, msg)
}

func (l *ExtendedLogger) SetLevel(level LogLevel) {
	l.level = level
}

func New(level LogLevel) *ExtendedLogger {
	return &ExtendedLogger{
		level:  level,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
