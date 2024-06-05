package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Cookie string `yaml:"Cookie"`
	Mongo  struct {
		Url      string `yaml:"Url"`
		Database string `yaml:"Database"`
	}
}

func Init(configPath string) {
	c := Config{}
	once.Do(func() {
		viper.SetConfigFile(configPath)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			err := viper.Unmarshal(&c)
			if err != nil {
				panic(err)
			}
			conf = &c
		})
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		err = viper.Unmarshal(&c)
		if err != nil {
			panic(err)
		}
		conf = &c
	})
}

func Get() *Config {
	return conf
}
