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
	ConsumerGroup     string `toml:"consumer_group"`
	ExtendedLogging   bool   `toml:"extended_logging"`
	InitialOffset     int64  `toml:"initial_offset"`
	CommitInterval    int    `toml:"commit_interval_sec"`
	MaxProcessingTime int    `toml:"max_processing_time_sec"`
	Brokers         []string
	Topics          []string
}

func NewConsumer(cfg *ConsumerConfig) (*cluster.Consumer, error) {
	if cfg.ExtendedLogging {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	conf := cluster.NewConfig()
	conf.ClientID = "butler{ .Project.Name }"

	if cfg.InitialOffset == sarama.OffsetNewest || cfg.InitialOffset == sarama.OffsetOldest {
		conf.Consumer.Offsets.Initial = cfg.InitialOffset
	} else {
		conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	if cfg.CommitInterval > 0 {
		conf.Consumer.Offsets.CommitInterval = time.Second * time.Duration(cfg.CommitInterval)
	} else {
		conf.Consumer.Offsets.CommitInterval = time.Minute
	}

	if cfg.MaxProcessingTime > 0 {
		conf.Consumer.MaxProcessingTime = time.Second * time.Duration(cfg.MaxProcessingTime)
	} else {
		conf.Consumer.MaxProcessingTime = time.Second * 30
	}

	conf.Consumer.Return.Errors = false
	conf.Group.Return.Notifications = false
	conf.MetricRegistry = metrics.NewPrefixedChildRegistry(metrics.DefaultRegistry, "sarama.")

	return cluster.NewConsumer(cfg.Brokers, cfg.ConsumerGroup, cfg.Topics, conf)
}
