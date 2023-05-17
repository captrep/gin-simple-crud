package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var Conf Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}

}
