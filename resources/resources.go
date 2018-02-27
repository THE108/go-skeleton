package resources

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/config"

	butler{if .Vars.useKafka}
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/resources/kafka"
	"github.com/bsm/sarama-cluster"
	butler{end}
)

type Resources struct {
	butler{if .Vars.useKafka}
	kafkaConsumer *cluster.Consumer
	butler{end}
}

func InitResources(cfg *config.Config) (res *Resources, err error) {
	res = new(Resources)

	butler{if .Vars.useKafka}
	res.kafkaConsumer, err = kafka.New(cfg)
	if err != nil {
		return nil, err
	}
	butler{end}

	return res, nil
}
