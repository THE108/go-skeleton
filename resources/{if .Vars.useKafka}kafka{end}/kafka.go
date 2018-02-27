package kafka

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/config"

	"stash.eyeota.com/tsuk/go-metrics/metrics"

	"github.com/bsm/sarama-cluster"
)

func New(cfg *config.Config) (*cluster.Consumer, error) {
	brokers, err := cfg.GetStrings(config.KafkaNodes)
	if err != nil {
		return nil, err
	}

	cgname, err := cfg.GetString(config.KafkaCounsumerGroup)
	if err != nil {
		return nil, err
	}

	topics, err := cfg.GetStrings(config.KafkaTopics)
	if err != nil {
		return nil, err
	}

	saramaConfig := cluster.NewConfig()
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Group.Return.Notifications = false
	saramaConfig.MetricRegistry = metrics.CreateChildRegistry("sarama.")

	return cluster.NewConsumer(brokers, cgname, topics, saramaConfig)
}
