package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	cfg, err := Parse("testdata/config.toml", "1.0.0", "1.10", "now", "git")
	if err != nil {
		t.Errorf("couldn't parse config.toml")
		return
	}

	counsumergroup, err := cfg.GetString("kafka.counsumergroup")
	if err != nil {
		t.Errorf("couldn't get string")
	}

	if counsumergroup != "consgrp" {
		t.Errorf("counsumergroup != 'consgrp'")
	}

	lvl1, err := cfg.GetString("tree.lvl1.key")
	if err != nil {
		t.Errorf("couldn't get string")
	}

	if lvl1 != "val1" {
		t.Errorf("lvl1 != 'val'")
	}

	lvl2, err := cfg.GetString("tree.lvl1.lvl2.key")
	if err != nil {
		t.Errorf("couldn't get string")
	}

	if lvl2 != "val2" {
		t.Errorf("lvl1 != 'val'")
	}

	kafkaLag, err := cfg.GetInt64("kafka.lag")
	if err != nil {
		t.Errorf("couldn't get int")
	}

	if kafkaLag != 123 {
		t.Errorf("kafkaLag != 123")
	}

	topics, err := cfg.GetStrings("kafka.topics")
	if err != nil {
		t.Errorf("couldn't get strings")
	}

	if len(topics) != 2 || topics[0] != "topic1" || topics[1] != "topic2" {
		t.Errorf("len(topics) != 2 || topics[0] != 'topic1' || topics[1] != 'topic2'")
	}
}
