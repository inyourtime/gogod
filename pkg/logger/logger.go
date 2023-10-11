package logger

import (
	"flag"
	"gogod/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logz *zap.Logger

// init initializes the logger configuration.
//
// No parameters.
// No return values.
func init() {
	dcCore := newDiscordCore()
	csCore := newConsoleCore()
	logz = zap.New(zapcore.NewTee(dcCore, csCore), zap.AddCaller(), zap.AddCallerSkip(1))

	defer logz.Sync()
}

func newDiscordCore() zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.StacktraceKey = ""

	return zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(&discordSink{}), zapcore.ErrorLevel)
}

func newConsoleCore() zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.StacktraceKey = ""
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(os.Stdout), zap.InfoLevel)
}

// Info logs a message with the "info" level.
//
// The message parameter is a string that represents the log message.
// The fields parameter is a variadic parameter that can accept zero or more zap.Field values.
// This function does not return anything.
func Info(message string, fields ...zap.Field) {
	logz.Info(message, fields...)
}

// Debug logs a debug message with optional fields.
//
// message: the debug message to log.
// fields: optional additional fields to include in the log.
func Debug(message string, fields ...zap.Field) {
	logz.Debug(message, fields...)
}

// Warn logs a warning message with additional fields.
//
// Parameters:
// - message: the warning message to be logged (string).
// - fields: additional fields to be included in the log (zap.Field).
func Warn(message string, fields ...zap.Field) {
	logz.Warn(message, fields...)
}

// Error logs an error message.
//
// The function takes a message parameter which can be either an error or a string.
// It also accepts an optional variadic parameter fields of type zap.Field.
// The function does not return any value.
func Error(message interface{}, fields ...zap.Field) {
	if flag.Lookup("test.v") == nil {
		switch v := message.(type) {
		case error:
			logz.Error(v.Error(), fields...)
		case string:
			logz.Error(v, fields...)
		}
	}
}

type discordSink struct{}

func (s *discordSink) Write(p []byte) (n int, err error) {
	if flag.Lookup("test.v") == nil {
		go WebhookSend(config.ENV.DiscordWebhook.ID, config.ENV.DiscordWebhook.Token, string(p))
	}
	return len(p), nil
}
