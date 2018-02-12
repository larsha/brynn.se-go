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
	Static     string
	Production bool
	Mailgun
}

func Get() Config {
	c := Config{
		Static:     os.Getenv("STATIC"),
		Production: os.Getenv("APP_ENV") == "production",
		Mailgun: Mailgun{
			Domain:       os.Getenv("MAILGUN_DOMAIN"),
			ApiKey:       os.Getenv("MAILGUN_API_KEY"),
			PublicApiKey: os.Getenv("MAILGUN_PUBLIC_API_KEY"),
			Email:        os.Getenv("MAILGUN_EMAIL"),
			Subject:      "Hello from fre.la!",
		},
	}

	return c
}
