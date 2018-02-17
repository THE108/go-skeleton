package config

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	data map[string]string
}

func Parse(filename, appVersion, goVersion, buildDate, gitLog string) (*Config, error) {
	cfg := &Config{
		data: make(map[string]string),
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

	return json.Unmarshal(content, &c.data)
}

func (c *Config) getValue(key string, parse func(string) (interface{}, error)) (interface{}, error) {
	value, found := c.data[key]
	if !found {
		return "", fmt.Errorf("can't find a key in config. key: %q", key)
	}

	return parse(value)
}

// GetBool returns config value by key as bool
func (c *Config) GetBool(key string, defaults ...bool) (bool, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return strconv.ParseBool(val)
	})

	var defaultValue bool
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(bool); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to bool")
}

// GetString returns config value by key as string
func (c *Config) GetString(key string, defaults ...string) (string, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return val, nil
	})

	var defaultValue string
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(string); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to string")
}

// GetInt returns config value by key as int
func (c *Config) GetInt(key string, defaults ...int) (int, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return strconv.Atoi(val)
	})

	var defaultValue int
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(int); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to int")
}

// GetInt64 returns config value by key as int
func (c *Config) GetInt64(key string, defaults ...int64) (int64, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return strconv.ParseInt(val, 10, 64)
	})

	var defaultValue int64
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(int64); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to int64")
}

// GetUint64 returns config value by key as uint64
func (c *Config) GetUint64(key string, defaults ...uint64) (uint64, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return strconv.ParseUint(val, 10, 64)
	})

	var defaultValue uint64
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(uint64); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to uint64")
}

// GetFloat64 returns config value by given key as float64
func (c *Config) GetFloat64(key string, defaults ...float64) (float64, error) {
	value, err := c.getValue(key, func(val string) (interface{}, error) {
		return strconv.ParseFloat(val, 64)
	})

	var defaultValue float64
	if len(defaults) > 0 {
		defaultValue = defaults[0]
	}

	if err != nil {
		return defaultValue, err
	}

	if v, ok := value.(float64); ok {
		return v, nil
	}

	return defaultValue, errors.New("can't cast interface to float64")
}
