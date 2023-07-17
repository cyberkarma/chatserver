package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"port"`
}

func LoadConfig() (Config, error) {
	var vp *viper.Viper
	vp = viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath(".")
	vp.SetConfigType("json")

	vp.SetDefault("port", "localhost:3001")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = vp.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil

}
