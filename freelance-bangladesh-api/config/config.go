package config

import (
	"bytes"
	_ "embed"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// ...

//go:embed config.yml
var defaultConfiguration []byte

var (
	configInstance *Config
	once           sync.Once
)

type Config struct {
	ListenIP   string
	ListenPort string
	Dashboard  *Dashboard
	Keycloak   *Keycloak
	Database   *Database
}

type Dashboard struct {
	Title string
	Realm string
}

type Keycloak struct {
	BaseUrl string
	Realm   string
	RestApi struct {
		ClientId     string
		ClientSecret string
	}
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigType("yml")
		viper.SetEnvPrefix("API")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		// Read configuration
		if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
			panic("Failed to read configuration: " + err.Error())
		}

		// Unmarshal the configuration
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			panic("Failed to unmarshal configuration: " + err.Error())
		}

		configInstance = &config
	})

	return configInstance
}
