package resources

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/config"

	butler{if .Vars.useKafka}
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/resources/kafka"
	butler{end}
)

type Resources struct {
	butler{if .Vars.useKafka}
	kafkaClient kafka.Broker
	butler{end}
}

func InitResources(cfg *config.Config) (*Resources, error) {
	res := &Resources{
		cfg: cfg,
	}

	butler{if .Vars.useKafka}
	res.kafkaClient, err = kafka.New(cfg)
	if err != nil {
		return nil, err
	}
	butler{end}

	return res, nil
}
