package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(path string) *Config {
	buffer, err := os.ReadFile(path)
	if err != nil {
		panic("error on read file: " + err.Error())
	}

	var cfg Config

	if err := yaml.Unmarshal(buffer, &cfg); err != nil {
		panic("error on unmarshal: " + err.Error())
	}

	return &cfg
}
