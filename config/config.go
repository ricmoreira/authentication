package config

import (
	"os"
)

var (
	HOST           string = "HOST"
	MONGO_HOST     string = "MONGO_HOST"
	MONGO_DATABASE string = "MONGO_DATABASE"
	COOKIE_DOMAIN  string = "COOKIE_DOMAIN"
)

type Config struct {
	Host              string
	MongoHost         string
	MongoDatabaseName string
	CookieDomain      string
}

func NewConfig() *Config {
	return &Config{
		Host:              MustGetEnv(HOST),
		MongoHost:         MustGetEnv(MONGO_HOST),
		MongoDatabaseName: MustGetEnv(MONGO_DATABASE),
		CookieDomain:      MustGetEnv(COOKIE_DOMAIN),
	}
}

func MustGetEnv(envVarName string) string {
	res, found := os.LookupEnv(envVarName)

	if !found {
		panic("Environment variable " + envVarName + " not found")
	}

	return res
}
