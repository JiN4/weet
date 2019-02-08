package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MySQL MySQL `envconfig:"MYSQL"`
}

type MySQL struct {
	Port     string
	User     string
	Pass     string
	Database string
}

func main() {
	var config Config
	envconfig.Process("", &config)
	fmt.Println(config.MySQL.Port)
}
