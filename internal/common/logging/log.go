package logging

import (
	"github.com/rs/zerolog"
	"io"
)

var Logger InternalLogger

func Fatal(msg string, args ...interface{}) {
	if len(args) > 0 {
		Logger.Logger.Fatal().Msgf(msg, args...)
	} else {
		Logger.Logger.Fatal().Msg(msg)
	}
}

func Crit(msg string, args ...interface{}) {
	if len(args) > 0 {
		Logger.Logger.Error().Msgf(msg, args...)
	} else {
		Logger.Logger.Error().Msg(msg)
	}
}

func Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		Logger.Logger.Info().Msgf(msg, args...)
	} else {
		Logger.Logger.Info().Msg(msg)
	}
}

func Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		Logger.Logger.Warn().Msgf(msg, args...)
	} else {
		Logger.Logger.Warn().Msg(msg)
	}
}

func Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		Logger.Logger.Debug().Msgf(msg, args...)
	} else {
		Logger.Logger.Debug().Msg(msg)
	}
}

func Init(dst io.Writer) {
	Logger = InternalLogger{
		zerolog.New(dst).With().Timestamp().Logger(),
	}
}

func SetLogLevel(level LogLevel) {
	zerolog.SetGlobalLevel(LogMapping[level])
}

type InternalLogger struct {
	zerolog.Logger
}

func (l *InternalLogger) GetSubLogger(k string, v string) *InternalLogger {
	return &InternalLogger{
		Logger: l.Logger.With().Str(k, v).Logger(),
	}
}

func (l *InternalLogger) Crit(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Error().Msgf(msg, args...)
	} else {
		l.Logger.Error().Msg(msg)
	}
}

func (l *InternalLogger) Fatal(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Fatal().Msgf(msg, args...)
	} else {
		l.Logger.Fatal().Msg(msg)
	}
}

func (l *InternalLogger) Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Warn().Msgf(msg, args...)
	} else {
		l.Logger.Warn().Msg(msg)
	}
}

func (l *InternalLogger) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Debug().Msgf(msg, args...)
	} else {
		l.Logger.Debug().Msg(msg)
	}
}

func (l *InternalLogger) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Logger.Info().Msgf(msg, args...)
	} else {
		l.Logger.Info().Msg(msg)
	}
}
