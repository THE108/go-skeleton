package kafka

import (
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/rcrowley/go-metrics"
)

type ConsumerConfig struct {
	ConsumerGroup   string `toml:"consumer_group"`
	ExtendedLogging bool   `toml:"extended_logging"`
	Brokers         []string
	Topics          []string
}

func NewConsumer(cfg *ConsumerConfig) (*cluster.Consumer, error) {
	if cfg.ExtendedLogging {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	conf := cluster.NewConfig()
	conf.ClientID = "butler{ .Project.Name }"
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Consumer.Offsets.CommitInterval = time.Minute
	conf.Consumer.MaxProcessingTime = 30 * time.Second
	conf.Consumer.Return.Errors = false
	conf.Group.Return.Notifications = false
	conf.MetricRegistry = metrics.NewPrefixedChildRegistry(metrics.DefaultRegistry, "sarama.")

	return cluster.NewConsumer(cfg.Brokers, cfg.ConsumerGroup, cfg.Topics, conf)
}
