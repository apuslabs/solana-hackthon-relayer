package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var configfile string

func init() {
	flag.StringVar(&configfile, "configfile", "", "yaml config path for server")
}

func Init() {
	flag.Parse()

	if configfile == "" {
		panic("configfile path must be set")
	}
	viper.SetConfigFile(configfile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		panic(err)
	}
}

func GetStr(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
