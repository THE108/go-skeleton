package app

import (
	"time"

	"github.com/THE108/go-skeleton/resources"

	"go.uber.org/zap/zapcore"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) doStaff() error {
	return nil
}

func (a *Application) stop() {

}

var a = NewApplication()

// Interfaces are standardised
type Config interface {
	GetString(key string, defaults ...string) (string, error)
	GetBool(key string, defaults ...bool) (bool, error)
	GetInt(key string, defaults ...int) (int, error)
	GetInt64(key string, defaults ...int64) (int64, error)
	GetUint64(key string, defaults ...uint64) (uint64, error)
	GetFloat64(key string, defaults ...float64) (float64, error)
}

type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsErrorEnabled() bool
}

type Monitoring interface {
	ObserveDuration(metricName string, startTime  time.Time)
}

// Hook called by skeleton
// Signature of that func is standardised
func Run(cfg Config, logger Logger, mon Monitoring, res *resources.Resources) error {
	// here user's code goes...
	return a.doStaff()
}

func Shutdown() {
	a.stop()
}
