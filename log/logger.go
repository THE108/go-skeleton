package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a logger which does nothing.
type Logger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

// NewLogger returns new Logger instance
func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: logger,
		sugaredLogger: logger.Sugar(),
	}, nil
}

// Structured logging (optimized)
func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

// Structured logging (empty interfaces)
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Infow(msg, keysAndValues...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Errorw(msg, keysAndValues...)
}

// Formatted logging
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *Logger) IsDebugEnabled() bool     { return l.levelEqualsTo(zapcore.DebugLevel) }
func (l *Logger) IsInfoEnabled() bool      { return l.levelEqualsTo(zapcore.InfoLevel) }
func (l *Logger) IsErrorEnabled() bool     { return l.levelEqualsTo(zapcore.ErrorLevel) }

func (l *Logger) levelEqualsTo(level zapcore.Level) bool {
	return l.logger.Core().Enabled(level)
}
