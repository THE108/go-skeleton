package cassandra

import (
	"time"

	"github.com/gocql/gocql"
)

type Config struct {
	Hosts    []string `toml:"hosts"`
	Keyspace string   `toml:"keyspace"`
}

func NewCluster(cfg *Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.One
	cluster.Timeout = time.Second * 3
	cluster.ReconnectInterval = time.Second * 10
	cluster.ConnectTimeout = time.Second * 5
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}
