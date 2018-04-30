package producer

import (
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rcrowley/go-metrics"
)

type Config struct {
	Brokers         []string
	ExtendedLogging bool `toml:"extended_logging"`
}

func NewProducer(cfg *Config) (sarama.SyncProducer, error) {
	if cfg.ExtendedLogging {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	producerConfig.ClientID = "butler{ .Project.Name }"
	producerConfig.Producer.Timeout = 30 * time.Second
	producerConfig.Producer.Return.Successes = true
	producerConfig.MetricRegistry = metrics.NewPrefixedChildRegistry(metrics.DefaultRegistry, "sarama.")

	return sarama.NewSyncProducer(cfg.Brokers, producerConfig)
}
