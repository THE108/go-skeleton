package config

const (
	AppVersion = "AppVersion"
	GoVersion  = "GoVersion"
	BuildDate  = "BuildDate"
	GitLog     = "GitLog"

	butler{if .Vars.useKafka}
	KafkaNodes = "KafkaNodes"
	KafkaCounsumerGroup = "KafkaCounsumerGroup"
	KafkaTopics = "KafkaTopics"
	butler{end}
)
