package kafka

import (
	"sync"
	"strings"

	"github.com/rcrowley/go-metrics"
)

type Monitoring struct {
	messagesPerTopicMu sync.RWMutex
	messagesPerTopic   map[string]metrics.Meter
}

func NewMonitoring() *Monitoring {
	return &Monitoring{
		messagesPerTopic: make(map[string]metrics.Meter),
	}
}

func (m *Monitoring) MarkMessage(topic string) {
	topic = strings.Replace(topic, ".", "_", -1)

	m.messagesPerTopicMu.RLock()
	meter, exists := m.messagesPerTopic[topic]
	m.messagesPerTopicMu.RUnlock()

	if !exists {
		meter = metrics.NewMeter()
		metrics.Register("topic."+topic, meter)

		m.messagesPerTopicMu.Lock()
		m.messagesPerTopic[topic] = meter
		m.messagesPerTopicMu.Unlock()
	}

	meter.Mark(1)
}
