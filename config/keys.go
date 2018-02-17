package config

const (
	AppVersion = "AppVersion"
	GoVersion  = "GoVersion"
	BuildDate  = "BuildDate"
	GitLog     = "GitLog"

	butler{if eq .Vars.useKafka}
	KafkaNodes = "KafkaNodes"
	butler{end}
)
