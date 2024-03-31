package config

import (
	"fmt"
	"os"
)

var (
	APPVERSION = getEnv("APPVERSION", "v0.1.0")
	// DB returns the name of the sqlite database
	DB = getEnv("DB", "host=192.168.0.6 user=gofinance password=gofinance dbname=gofinance port=5432 sslmode=disable TimeZone=America/Sao_Paulo")
	// PORT returns the server listening port
	PORT = getEnv("PORT", "4000")
	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "secretKey")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP        = getEnv("TOKEN_EXP", "10h")
	APP_ENVIRONMENT = getEnv("APP_ENVIRONMENT", "development")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
