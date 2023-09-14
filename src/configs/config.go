package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server int    `mapstructure:"server"`
	DB     string `mapstructure:"db"`
}

func LoadConfig() (Config, error) {
	var vp *viper.Viper
	vp = viper.New()
	var config Config
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
