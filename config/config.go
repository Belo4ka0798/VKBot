package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	VKBot VKBotConfig `yaml:"VKBot"`
}

type VKBotConfig struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

func ReadConfig(filename string) (*Config, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
