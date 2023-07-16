package configs

import (
	"github.com/spf13/viper"
)

type ProdConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Pc ProdConfig `mapstructure:"prod"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath(".")
	vp.SetConfigType("json")

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
