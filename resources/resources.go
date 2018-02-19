package resources

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/config"

	butler{if .Vars.useKafka}
	"strings"
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/resources/kafka"
	butler{end}
)

type Resources struct {
	cfg *config.Config

	butler{if .Vars.useKafka}
	kafkaClient kafka.Broker
	butler{end}
}

func InitResources(cfg *config.Config) (*Resources, error) {
	res := &Resources{
		cfg: cfg,
	}

	butler{if .Vars.useKafka}
	nodesString, err := cfg.GetString(config.KafkaNodes)
	if err != nil {
		return nil, err
	}

	res.kafkaClient, err = kafka.NewBroker(strings.Split(nodesString, ","))
	butler{end}

	return res, nil
}
