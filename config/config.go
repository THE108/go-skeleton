package config

import (
	"fmt"

	"butler{ .Vars.repoPath }/butler{ .Project.Name }/monitoring"
	butler{if .Vars.useKafka}
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/kafka"
	butler{end}

	"github.com/BurntSushi/toml"
)

type Config struct {
	butler{if .Vars.useKafka}
	KafkaConsumer kafka.ConsumerConfig `toml:"kafka_consumer"`
	KafkaProducer kafka.ProducerConfig `toml:"kafka_producer"`
	butler{end}

	Monitoring monitoring.Config
}

func (cfg *Config) applyDefaultsAndValidate() error {
	butler{if .Vars.useKafka}
	if cfg.KafkaConsumer.ConsumerGroup == "" {
		return fmt.Errorf("kafka consumer group must be not empty")
	}

	if len(cfg.KafkaConsumer.Topics) == 0 {
		return fmt.Errorf("kafka topics list must be not empty")
	}

	if len(cfg.KafkaConsumer.Brokers) == 0 {
		return fmt.Errorf("kafka brockers list must be not empty")
	}

	if len(cfg.KafkaProducer.Brokers) == 0 {
		return fmt.Errorf("kafka brockers list must be not empty")
	}
	butler{end}

	return nil
}

func Parse(file string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return nil, err
	}

	if err := conf.applyDefaultsAndValidate(); err != nil {
		return nil, err
	}

	return &conf, nil
}
