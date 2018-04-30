package abstract

import (
	"time"

	"github.com/rcrowley/go-metrics"
)

type Monitoring struct {
	requestRate        metrics.Meter
	requestErrorRate   metrics.Meter
	requestLatency     metrics.Timer
}

func NewMonitoring() *Monitoring {
	m := &Monitoring{
		requestRate:      metrics.NewMeter(),
		requestErrorRate: metrics.NewMeter(),
		requestLatency:   metrics.NewTimer(),
	}

	metrics.Register("request_rate", m.requestRate)
	metrics.Register("request_error_rate", m.requestErrorRate)
	metrics.Register("request_latency", m.requestLatency)

	return m
}

func (m *Monitoring) MarkRequest(success bool, elapsed time.Duration) {
	if success {
		m.requestRate.Mark(1)
	} else {
		m.requestErrorRate.Mark(1)
	}

	m.requestLatency.Update(elapsed)
}
