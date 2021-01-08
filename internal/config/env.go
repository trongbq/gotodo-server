package config

import (
	"fmt"
	"os"
)

const (
	LocalEnv = "local"
	TestEnv  = "test"
	AlphaEnv = "alpha"
	BetaEnv  = "beta"
	ProdEnv  = "prod"
)

func GetEnv(key, fallback string) string {
	if s, ok := os.LookupEnv(key); ok {
		return s
	}
	return fallback
}

func GetEnvStrict(key string) string {
	if s, ok := os.LookupEnv(key); ok {
		return s
	}
	panic(fmt.Sprintf("ENV %v is missing, please add it", key))
}

func GetEnvInt(key, fallback string) string {
	if s, ok := os.LookupEnv(key); ok {
		return s
	}
	return fallback
}
