package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func AppConfig() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	port := viper.GetString("prod.port")
	fmt.Println(port, "port is init here")

}
