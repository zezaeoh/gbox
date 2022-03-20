package logger

import (
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LogLevel uint8

var logger *zap.Logger

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "%",
	zapcore.InfoLevel:   "*",
	zapcore.WarnLevel:   "-",
	zapcore.ErrorLevel:  "!",
	zapcore.DPanicLevel: "!",
	zapcore.PanicLevel:  "!!",
	zapcore.FatalLevel:  "!!",
}

var logLevelColor = map[zapcore.Level]func(a ...interface{}) string{
	zapcore.DebugLevel:  color.New().SprintFunc(),
	zapcore.InfoLevel:   color.New(color.FgCyan).SprintFunc(),
	zapcore.WarnLevel:   color.New(color.FgYellow).SprintFunc(),
	zapcore.ErrorLevel:  color.New(color.FgHiRed).SprintFunc(),
	zapcore.DPanicLevel: color.New(color.FgWhite, color.BgRed).SprintFunc(),
	zapcore.PanicLevel:  color.New(color.FgWhite, color.BgRed).SprintFunc(),
	zapcore.FatalLevel:  color.New(color.FgWhite, color.BgRed).SprintFunc(),
}

func init() {
	config := zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "severity",
		CallerKey:        "caller",
		EncodeLevel:      customEncodeLevel,
		EncodeCaller:     zapcore.FullCallerEncoder,
		ConsoleSeparator: " ",
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	level := zap.InfoLevel
	if os.Getenv("GBOX_DEBUG") == "true" {
		level = zap.DebugLevel
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), consoleDebugging, level),
	)

	logger = zap.New(core)
}

func customEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelColor[level](logLevelSeverity[level]))
}

func Logger() *zap.SugaredLogger {
	return logger.Sugar()
}
