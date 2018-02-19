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

func (m *Monitoring) Mark(metricName string, value int64) {
	metrics.GetOrRegisterMeter(metricName, nil).Mark(value)
}

func (m *Monitoring) UpdateTimer(metricName string, duration time.Duration) {
	metrics.GetOrRegisterHistogram(metricName, nil, m.sample).Update(duration.Nanoseconds() / 1e6)
}

func (m *Monitoring) UpdateGauge(metricName string, value int64) {
	metrics.GetOrRegisterGauge(metricName, nil).Update(value)
}

func (m *Monitoring) GetMetric(metricName string) interface{}  {
	return metrics.Get(metricName)
}

func (m *Monitoring) DeregisterMetric(metricName string) {
}