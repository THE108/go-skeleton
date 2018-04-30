package config

import (
	"fmt"

	"butler{ .Project.Name }/butler{ .Project.Name }/monitoring"
	butler{if .Vars.useKafkaConsumer}
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/kafka/consumer"
	butler{end}
	butler{if .Vars.useKafkaProducer}
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/kafka/producer"
	butler{end}

	"github.com/BurntSushi/toml"
)

type Config struct {
	butler{if .Vars.useKafkaConsumer}
	KafkaConsumer consumer.Config `toml:"kafka_consumer"`
	butler{end}
	butler{if .Vars.useKafkaProducer}
	KafkaProducer producer.Config `toml:"kafka_producer"`
	butler{end}

	Monitoring monitoring.Config
}

func (cfg *Config) applyDefaultsAndValidate() error {
	butler{if .Vars.useKafkaConsumer}
	if cfg.KafkaConsumer.ConsumerGroup == "" {
		return fmt.Errorf("kafka consumer group must be not empty")
	}

	if len(cfg.KafkaConsumer.Topics) == 0 {
		return fmt.Errorf("kafka topics list must be not empty")
	}

	if len(cfg.KafkaConsumer.Brokers) == 0 {
		return fmt.Errorf("kafka brockers list must be not empty")
	}
	butler{end}

	butler{if .Vars.useKafkaProducer}
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
