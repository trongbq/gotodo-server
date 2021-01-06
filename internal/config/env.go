package config

import "os"

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

func GetEnvInt(key, fallback string) string {
	if s, ok := os.LookupEnv(key); ok {
		return s
	}
	return fallback
}
