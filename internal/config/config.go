package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB       DBConfig       `yaml:"db"`
	Bot      BotConfig      `yaml:"bot"`
	Notifier NotifierConfig `yaml:"notifier"`
}

type DBConfig struct {
	Name string `yaml:"name"`
}

type BotConfig struct {
	Token string `yaml:"token"`
}

type NotifierConfig struct {
	SendInterval int `yaml:"sendInterval"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
