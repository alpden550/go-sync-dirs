package logging

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"log"
)

type Logger struct{ *logrus.Logger }

func NewLogger() *Logger {
	logLevel := logrus.InfoLevel
	formatter := &logrus.TextFormatter{
		PadLevelText:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.SetFormatter(formatter)

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "log.txt",
		MaxSize:    10, // megabytes
		MaxBackups: 3,  // amounts
		MaxAge:     30, //days
		Level:      logLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	logger.AddHook(rotateFileHook)

	return &Logger{logger}
}

func LoggerFromContext(ctx context.Context, name string) *Logger {
	if l, ok := ctx.Value(name).(*Logger); ok {
		return l
	}
	return NewLogger()
}
