package main

import (
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/config"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()

	title := viper.Get("Dashboard.Title")
	fmt.Println(title, "api is running...")
}
