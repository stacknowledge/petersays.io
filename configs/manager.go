package configs

import "github.com/jinzhu/configor"

type (
	Config struct {
		Environment string `required:"true"`

		Application struct {
			Name        string `required:"true"`
			Description string `required:"true"`
			Address     string `required:"true"`
			Port        string `required:"true"`
		}
	}
)

func NewConfigurations() *Config {
	configurations := &Config{}
	return configurations.Load()
}

func (config *Config) Load() *Config {
	configuration := new(Config)
	configor.Load(configuration, "./configs/config.yml")
	return configuration
}
