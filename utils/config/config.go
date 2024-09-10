package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	PgConnStr string `env:"PG_CONNECTION_STRING"`
	AppEnv    string `env:"APP_ENV"`
}

func LoadConfig() (Config, error) {

	config := Config{}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// or you can use generics
	config, err := env.ParseAs[Config]()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return config, err
}
