package config

import (
	"os"
)

func SetEnvironmentVariables() {

	os.Setenv("LOG_FILE", "log.txt")
	os.Setenv("URI", "mongodb://localhost:27017")
	os.Setenv("DB_NAME", "data-gin")
	os.Setenv("COLLECTION_NAME", "userDetails")
	os.Setenv("COLLECTION_NAME_AUTH", "AuthenticationUserDetails")
	os.Setenv("PORT", "localhost:8008")
}
