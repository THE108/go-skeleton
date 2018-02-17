package monitoring

import (
	"time"

	"github.com/rcrowley/go-metrics"
)

type Monitoring struct {
	sample metrics.Sample
}

func NewMonitoring() *Monitoring {
	return &Monitoring{
		sample: metrics.NewExpDecaySample(1028, 0.015),
	}
}

func (m *Monitoring) ObserveDuration(metricName string, startTime  time.Time) {
	metrics.GetOrRegisterHistogram(metricName, nil, m.sample).Update(time.Since(startTime).Nanoseconds() / 1e6)
}
