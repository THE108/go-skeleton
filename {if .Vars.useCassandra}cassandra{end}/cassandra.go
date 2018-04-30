package cassandra

import (
	"time"

	"github.com/gocql/gocql"
)

type Config struct {
	CassandraHosts    []string `toml:"cassandra_hosts"`
	CassandraKeyspace string   `toml:"cassandra_keyspace"`
}

func NewCluster(cfg *Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(cfg.CassandraHosts...)
	cluster.Keyspace = cfg.CassandraKeyspace
	cluster.Consistency = gocql.One
	cluster.Timeout = time.Second * 3
	cluster.ReconnectInterval = time.Second * 10
	cluster.ConnectTimeout = time.Second * 5
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
}
