package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(("api"))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize config with viper: %w", err))
	}

	fmt.Println("initialized config successfully with viper!")
}
