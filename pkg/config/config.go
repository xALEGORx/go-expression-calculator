package config

import "os"

type iconfig struct {
	ServerAddr string
	Mode       string

	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string

	RabbitURL string
}

var config *iconfig

func Init() *iconfig {
	config = &iconfig{
		ServerAddr: getEnv("SERVER_ADDR", "localhost:8080"),
		Mode:       getEnv("MODE", "debug"),

		PostgresUser:     getEnv("POSTGRES_USER", "calculator"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "test12345"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDatabase: getEnv("POSTGRES_DATABASE", "calculator"),

		RabbitURL: getEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672"),
	}

	return config
}

func Get() *iconfig {
	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
