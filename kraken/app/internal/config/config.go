package config

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var instance *Config
var once sync.Once

type Kraken struct {
	Key string `env:"KRAKEN_API_KEY"       env-required:""`
}

type Logger struct {
	DisableCaller     bool   `yaml:"disable_caller"`
	Development       bool   `yaml:"development"`
	DisableStacktrace bool   `yaml:"disable_stacktrace"`
	Encoding          string `yaml:"encoding"`
	Level             string `yaml:"level"`
}

type Config struct {
	Logger Logger `yaml:"logger"`
	Kraken Kraken
}

func GetConfig() (*Config, error) {
	var cfgErr error
	once.Do(func() {
		instance = &Config{}
		dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
		if exist, _ := Exists(containerConfigPath + "/configs/config.yaml"); exist {
			if err := cleanenv.ReadConfig(containerConfigPath+"/configs/config.yaml", instance); err != nil {
				cfgErr = err
			}
		}
	})
	return instance, cfgErr
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
