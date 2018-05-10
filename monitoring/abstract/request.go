package abstract

import (
	"time"

	"github.com/rcrowley/go-metrics"
)

type Monitoring struct {
	rate      metrics.Meter
	errorRate metrics.Meter
	latency   metrics.Timer
}

func NewMonitoring(prefix string) *Monitoring {
	m := &Monitoring{
		rate:      metrics.NewMeter(),
		errorRate: metrics.NewMeter(),
		latency:   metrics.NewTimer(),
	}

	metrics.Register(prefix + "_rate", m.rate)
	metrics.Register(prefix + "_error_rate", m.errorRate)
	metrics.Register(prefix + "_latency", m.latency)

	return m
}

func (m *Monitoring) MarkRequest(success bool, elapsed time.Duration) {
	if success {
		m.rate.Mark(1)
	} else {
		m.errorRate.Mark(1)
	}

	m.latency.Update(elapsed)
}
