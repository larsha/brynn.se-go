package config

import (
	"os"
)

type Mailgun struct {
	Domain       string
	ApiKey       string
	PublicApiKey string
	Email        string
	Subject      string
}

type Config struct {
	StaticFolder string
	Production   bool
	Cachebust    string
	Mailgun
}

func Get() *Config {
	c := &Config{
		StaticFolder: "/static",
		Production:   os.Getenv("APP_ENV") == "production",
		Cachebust:    os.Getenv("CACHEBUST"),
		Mailgun: Mailgun{
			Domain:       os.Getenv("MAILGUN_DOMAIN"),
			ApiKey:       os.Getenv("MAILGUN_API_KEY"),
			PublicApiKey: os.Getenv("MAILGUN_PUBLIC_API_KEY"),
			Email:        os.Getenv("MAILGUN_EMAIL"),
			Subject:      "Hello from brynn.se!",
		},
	}

	return c
}
