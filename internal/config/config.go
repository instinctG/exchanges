package config

import "github.com/spf13/viper"

type Config struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	LogLevel string `mapstructure:"loglevel"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
