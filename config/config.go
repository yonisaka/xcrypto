package config

import (
	"os"
	"strconv"
)

type (
	Config struct {
		AppName string
		AppPort int
		RSA     RSAConfig
	}

	CacheConfig struct {
		Addr     string
		Port     int
		Password string
		DB       int
	}

	DBConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	RSAConfig struct {
		PublicKeyPath  string
		PrivateKeyPath string
	}
)

func New() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "go-xcrypto"),
		AppPort: getEnvAsInt("APP_PORT", 8080),
		RSA: RSAConfig{
			PublicKeyPath:  getEnv("RSA_PUBLIC_KEY", ""),
			PrivateKeyPath: getEnv("RSA_PRIVATE_KEY", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}
