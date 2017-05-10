package config

import (
	"os"
)

type Config struct {
	Production bool
	Cachebust  string
}

func Get() *Config {
	c := &Config{
		Production: os.Getenv("APP_ENV") == "production",
		Cachebust:  os.Getenv("CACHEBUST"),
	}

	return c
}
