package config

import (
	"fmt"

	"butler{ .Vars.repoPath }/butler{ .Project.Name }/monitoring"

	butler{if .Vars.useKafka}
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/kafka"
	butler{end}

	butler{if .Vars.useCassandra}
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/cassandra"
	butler{end}

	"github.com/BurntSushi/toml"
)

type Config struct {
	Monitoring monitoring.Config

	butler{if .Vars.useKafka}
	KafkaConsumer kafka.ConsumerConfig `toml:"kafka_consumer"`
	KafkaProducer kafka.ProducerConfig `toml:"kafka_producer"`
	butler{end}

	butler{if .Vars.useCassandra}
	Cassandra cassandra.Config `toml:"cassandra"`
	butler{end}
}

func (cfg *Config) applyDefaultsAndValidate() error {
	if cfg.Monitoring.Server == "" {
		return fmt.Errorf("monitoring server must be not empty")
	}

	if cfg.Monitoring.Prefix == "" {
		return fmt.Errorf("monitoring prefix must be not empty")
	}

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

	butler{if .Vars.useCassandra}
	if len(cfg.Cassandra.Hosts) == 0 {
		return fmt.Errorf("cassandra hosts list must be not empty")
	}

	if cfg.Cassandra.Keyspace == "" {
		return fmt.Errorf("cassandra keyspace must be not empty")
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
