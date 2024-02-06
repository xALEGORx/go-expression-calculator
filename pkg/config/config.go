package config

import (
	"github.com/joho/godotenv"
	"os"
)

type IConfig struct {
	ServerAddr string
	Mode       string

	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string

	RabbitURL   string
	RabbitQueue string
}

var config *IConfig

func Init() *IConfig {
	godotenv.Load()

	config = &IConfig{
		ServerAddr: getEnv("SERVER_ADDR", "localhost:8080"),
		Mode:       getEnv("MODE", "debug"),

		PostgresUser:     getEnv("POSTGRES_USER", "calculator"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "test12345"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDatabase: getEnv("POSTGRES_DATABASE", "calculator"),

		RabbitURL:   getEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672"),
		RabbitQueue: getEnv("RABBIT_QUEUE", "CalculatorTaskQueue1"),
	}

	return config
}

func Get() *IConfig {
	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
