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
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.PanicLevel),
		//zap.Fields(),
		//zap.Hooks(),
	)

	if err != nil {
		return nil, err
	}

	return &Logger{
		logger:        logger,
		sugaredLogger: logger.Sugar(),
	}, nil
}

func (l *Logger) Close() error {
	return l.logger.Sync()
}

// Structured logging (empty interfaces)
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Infow(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Errorw(msg, keysAndValues...)
}

func (l *Logger) IsDebugEnabled() bool { return l.logger.Core().Enabled(zapcore.DebugLevel) }
func (l *Logger) IsInfoEnabled() bool  { return l.logger.Core().Enabled(zapcore.InfoLevel) }
func (l *Logger) IsErrorEnabled() bool { return l.logger.Core().Enabled(zapcore.ErrorLevel) }
