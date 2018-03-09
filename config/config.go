package config

import (
	"errors"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	ErrorNotFound     = errors.New("config key not found")
	ErrorTypeMismatch = errors.New("config value type mismatch")
)

type Config struct {
	data map[string]interface{}
}

func Parse(filename, appVersion, goVersion, buildDate, gitLog string) (*Config, error) {
	cfg := &Config{
		data: make(map[string]interface{}),
	}

	if err := cfg.parseConfigFile(filename); err != nil {
		return nil, err
	}

	cfg.data[AppVersion] = appVersion
	cfg.data[GoVersion] = goVersion
	cfg.data[BuildDate] = buildDate
	cfg.data[GitLog] = gitLog

	return cfg, nil
}

func (c *Config) parseConfigFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(filename, "json"):
		err = json.Unmarshal(content, &c.data)
	case strings.HasSuffix(filename, "toml"):
		_, err = toml.Decode(string(content), &c.data)
	}

	return err
}

func (c *Config) getValue(key string) (interface{}, bool) {
	return getValueTree(strings.Split(key, "."), c.data)
}

func getValueTree(tree []string, data map[string]interface{}) (interface{}, bool) {
	if len(tree) == 0 {
		return nil, false
	}

	value, found := data[tree[0]]
	if !found {
		return nil, false
	}

	if len(tree) == 1 {
		return value, true
	}

	if mapValue, ok := value.(map[string]interface{}); ok {
		return getValueTree(tree[1:], mapValue)
	}

	return nil, false
}

butler{ range .Vars.types }
type T{{.}} {{.}}
{end}

func (c *Config) GetStrings(key string, defaults ...string) (result []string, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.([]string); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}

// GetInt returns config value by key as int
func (c *Config) GetInt(key string, defaults ...int) (result int, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults[0]
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.(int); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}

// GetString returns config value by key as string
func (c *Config) GetString(key string, defaults ...string) (result string, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults[0]
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.(string); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}

// GetInt64 returns config value by key as int64
func (c *Config) GetInt64(key string, defaults ...int64) (result int64, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults[0]
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.(int64); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}

// GetUint64 returns config value by key as uint64
func (c *Config) GetUint64(key string, defaults ...uint64) (result uint64, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults[0]
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.(uint64); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}

// GetFloat64 returns config value by key as float64
func (c *Config) GetFloat64(key string, defaults ...float64) (result float64, err error) {
	value, found := c.getValue(key)
	if !found {
		if len(defaults) > 0 {
			result = defaults[0]
			return
		}

		err = ErrorNotFound
		return
	}

	if v, ok := value.(float64); ok {
		result = v
		return
	}

	err = ErrorTypeMismatch
	return
}
