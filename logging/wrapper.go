package logging

import (
	"go.uber.org/zap/zapcore"
)

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}
func DPanic(msg string, fields ...zapcore.Field) {
	logger.DPanic(msg, fields...)
}
func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

func Debugf(msg string, fields ...any) {
	logger.Sugar().Debugf(msg, fields...)
}
func Infof(msg string, fields ...any) {
	logger.Sugar().Infof(msg, fields...)
}
func Warnf(msg string, fields ...any) {
	logger.Sugar().Warnf(msg, fields...)
}
func Errorf(msg string, fields ...any) {
	logger.Sugar().Errorf(msg, fields...)
}
func DPanicf(msg string, fields ...any) {
	logger.Sugar().DPanicf(msg, fields...)
}
func Fatalf(msg string, fields ...any) {
	logger.Sugar().Fatalf(msg, fields...)
}