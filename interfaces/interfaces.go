// Standardised interfaces
package interfaces

import "time"

type Config interface {
	GetString(key string, defaults ...string) (string, error)
	GetBool(key string, defaults ...bool) (bool, error)
	GetInt(key string, defaults ...int) (int, error)
	GetInt64(key string, defaults ...int64) (int64, error)
	GetUint64(key string, defaults ...uint64) (uint64, error)
	GetFloat64(key string, defaults ...float64) (float64, error)
}

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsErrorEnabled() bool
}

type Monitoring interface {
	Mark(metricName string, value int64)
	UpdateTimer(metricName string, duration time.Duration)
	UpdateGauge(metricName string, value int64)
	GetMetric(metricName string) interface{}
	DeregisterMetric(metricName string)
}
