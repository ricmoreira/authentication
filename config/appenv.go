package appenv

import (
	"os"
)

var (
	HOST           string = "HOST"
	MONGO_HOST     string = "MONGO_HOST"
	MONGO_DATABASE string = "MONGO_DATABASE"
	COOKIE_DOMAIN  string = "COOKIE_DOMAIN"
)

func CheckMandatoryEnv() {
	MustGetEnv(HOST)
	MustGetEnv(MONGO_HOST)
	MustGetEnv(MONGO_DATABASE)
	MustGetEnv(COOKIE_DOMAIN)
}

func MustGetEnv(envVarName string) string {
	res, found := os.LookupEnv(envVarName)

	if !found {
		panic("Environment variable " + envVarName + " not found")
	}

	return res
}
