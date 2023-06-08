package plugin

import (
	"fmt"
	"strings"
)

func logPrefix(p Plugin) string {
	return fmt.Sprintf("plugin %s: ", strings.ToLower(p.Name()))
}

// LogDebugf logs a debug message with the common log prefix.
func LogDebugf(p Plugin, format string, args ...any) {
	log.Debugf(logPrefix(p)+format, args...)
}

// LogInfof logs an info message with the common log prefix.
func LogInfof(p Plugin, format string, args ...any) {
	log.Infof(logPrefix(p)+format, args...)
}

// LogWarnf logs a warning message with the common log prefix.
func LogWarnf(p Plugin, format string, args ...any) {
	log.Warnf(logPrefix(p)+format, args...)
}
