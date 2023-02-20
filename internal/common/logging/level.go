package logging

import (
	"github.com/rs/zerolog"
)

var LogMapping = map[LogLevel]zerolog.Level{
	LogLevel_DEBUG:                zerolog.DebugLevel,
	LogLevel_INFO:                 zerolog.InfoLevel,
	LogLevel_WARNING:              zerolog.WarnLevel,
	LogLevel_CRITICAL:             zerolog.FatalLevel,
	LogLevel_LOGLEVEL_UNSPECIFIED: zerolog.Disabled,
}

func GetLogLevelFromString(s string) LogLevel {

	if level, ok := LogLevel_value[s]; ok {
		return LogLevel(level)
	}
	return LogLevel_LOGLEVEL_UNSPECIFIED
}
