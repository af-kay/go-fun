package extendedlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixes(t *testing.T) {
	assert.Equal(t, len(prefixes), 5)

	assert.Equal(t, "ERROR", prefixes[LogLevelError])
	assert.Equal(t, "WARN", prefixes[LogLevelWarning])
	assert.Equal(t, "INFO", prefixes[LogLevelInfo])
	assert.Equal(t, "DEBUG", prefixes[LogLevelDebug])
	assert.Equal(t, "TRACE", prefixes[LogLevelTrace])
}

func TestExtendedLogger(t *testing.T) {
	infoLogger := New(LogLevelInfo)

	infoLogger.Errorln("Error - OK")
	infoLogger.Warnln("Warning - OK")
	infoLogger.Infoln("Info - OK")
	infoLogger.Debugln("Info - FAIL")
	infoLogger.Traceln("Trace - FAIL")

	infoLogger.SetLevel(LogLevelTrace)
	infoLogger.Debugln("Debug - OK")
	infoLogger.Traceln("Trace - OK")

}
