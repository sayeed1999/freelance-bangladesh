package config

import (
	"bytes"
	_ "embed"
	"strings"

	"github.com/spf13/viper"
)

// ...

//go:embed config.yml
var defaultConfiguration []byte

type Config struct {
	ListenIP   string
	ListenPort string
	Dashboard  *Dashboard
	Keycloak   *Keycloak
}

type Dashboard struct {
	Title string
	Realm string
}

type Keycloak struct {
	BaseUrl string
	Realm   string
	RestApi *RestApi
}

type RestApi struct {
	ClientId     string
	ClientSecret string
}

type Postgres struct {
	Host     string
	User     string
	Password string
}

func Read() (*Config, error) {
	viper.SetConfigType("yml")
	viper.SetEnvPrefix("API")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Configuration file
	viper.SetConfigType("yml")

	// Read configuration
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return nil, err
	}

	// Unmarshal the configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
