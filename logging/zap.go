package logging

import (
	"log"
	"os"

	"github.com/sslime336/bot-486/config"
	"github.com/sslime336/bot-486/global/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Init() {
	localLogPath := config.App.Log.Path
	initLogPath(localLogPath)

	encoderConf := zap.NewDevelopmentEncoderConfig()
	generalLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.DebugLevel
	})
	if os.Getenv("AIRP_MODE") == "release" {
		encoderConf = zap.NewProductionEncoderConfig()
		generalLevel = zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev >= zap.InfoLevel
		})
	}
	prepareEncoderConf(&encoderConf)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(os.Stdout),
		generalLevel,
	)

	generalSyncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/bot.log",
		MaxSize:    1024, // 1GiB
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
	})
	errorSyncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    500,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
	})

	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder

	generalCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(generalSyncWriter),
		generalLevel,
	)
	errorCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(errorSyncWriter),
		zap.NewAtomicLevelAt(zap.ErrorLevel),
	)

	logger = zap.New(zapcore.NewTee(consoleCore, generalCore, errorCore),
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	)
}

func Logger() *zap.Logger {
	return logger
}

func initLogPath(path string) {
	if exist, err := utils.PathExists(path); exist {
		return
	} else if err != nil {
		log.Fatal("could not init log path", err)
	}
	if err := os.Mkdir(path, 0644); err != nil {
		log.Fatal("failed to make log path", err)
	}
}

func prepareEncoderConf(encoderConfig *zapcore.EncoderConfig) {
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "name"
	encoderConfig.CallerKey = "caller"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
}
