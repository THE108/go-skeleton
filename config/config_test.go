package config

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type configSuite struct{}

var _ = Suite(&configSuite{})

func (s configSuite) Test(c *C) {
	cfg, err := Parse("testdata/config.toml")

	c.Assert(err, IsNil)

	butler{if .Vars.useKafkaConsumer}
	c.Assert(cfg.KafkaConsumer.ConsumerGroup, Equals, "kafkaproxy")
	c.Assert(cfg.KafkaConsumer.Topics, DeepEquals, []string{"topic1", "topic2"})
	c.Assert(cfg.KafkaConsumer.Brokers, DeepEquals, []string{"localhost:9092"})
	butler{end}

	butler{if .Vars.useKafkaProducer}
	c.Assert(cfg.KafkaProducer.Brokers, DeepEquals, []string{"localhost:9092"})
	butler{end}

	c.Assert(cfg.Monitoring.Server, Equals, "localhost:2005")
	c.Assert(cfg.Monitoring.Prefix, Equals, "test")
}
