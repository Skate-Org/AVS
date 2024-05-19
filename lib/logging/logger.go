package logging

import (
	"io"
	"os"

	zerolog "github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(outputWriter io.Writer) *Logger {
	logger := zerolog.New(outputWriter).With().Timestamp().Logger()
	return &Logger{
		logger: logger,
	}
}

func NewLoggerWithConsoleWriter() *Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	return NewLogger(consoleWriter)
}

func (z Logger) Debug(msg string, tags ...interface{}) {
	z.logger.Debug().Fields(parseTags(tags...)).Msg(msg)
}

func (z Logger) Info(msg string, tags ...interface{}) {
	z.logger.Info().Fields(parseTags(tags...)).Msg(msg)
}

func (z Logger) Warn(msg string, tags ...interface{}) {
	z.logger.Warn().Fields(parseTags(tags...)).Msg(msg)
}

func (z Logger) Error(msg string, tags ...interface{}) {
	z.logger.Error().Fields(parseTags(tags...)).Msg(msg)
}

func (z Logger) Fatal(msg string, tags ...interface{}) {
	z.logger.Fatal().Fields(parseTags(tags...)).Msg(msg)
	os.Exit(1)
}

func (z Logger) Debugf(template string, args ...interface{}) {
	z.logger.Debug().Msgf(template, args...)
}

func (z Logger) Infof(template string, args ...interface{}) {
	z.logger.Info().Msgf(template, args...)
}

func (z Logger) Warnf(template string, args ...interface{}) {
	z.logger.Warn().Msgf(template, args...)
}

func (z Logger) Errorf(template string, args ...interface{}) {
	z.logger.Error().Msgf(template, args...)
}

func (z Logger) Fatalf(template string, args ...interface{}) {
	z.logger.Fatal().Msgf(template, args...)
	os.Exit(1)
}

func (z *Logger) With(tags ...interface{}) *Logger {
	if len(tags) == 0 {
		return z
	}
	fields := parseTags(tags...)
	return &Logger{
		logger: z.logger.With().Fields(fields).Logger(),
	}
}

func parseTags(tags ...interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	for i := 0; i < len(tags); i += 2 {
		if i+1 < len(tags) {
			key, ok := tags[i].(string)
			if !ok {
				// Skip if the key is not a string
				continue
			}
			fields[key] = tags[i+1]
		}
	}
	return fields
}
