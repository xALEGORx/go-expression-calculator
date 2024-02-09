package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type IConfig struct {
	ServerAddr string
	Mode       string

	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string

	RabbitURL        string
	RabbitTaskQueue  string
	RabbitAgentQueue string

	AgentTimeout     int
	AgentPing        int
	AgentResolveTime int
}

var config *IConfig

func Init() (*IConfig, error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	config = &IConfig{
		ServerAddr: getEnv("SERVER_ADDR", "localhost:8080"),
		Mode:       getEnv("MODE", "debug"),

		PostgresUser:     getEnv("POSTGRES_USER", "calculator"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "test12345"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDatabase: getEnv("POSTGRES_DATABASE", "calculator"),

		RabbitURL:        getEnv("RABBIT_URL", "amqp://guest:guest@localhost:5672"),
		RabbitTaskQueue:  getEnv("RABBIT_TASK_QUEUE", "CalculatorTaskQueue1"),
		RabbitAgentQueue: getEnv("RABBIT_AGENT_QUEUE", "CalculatorAgentQueue1"),
	}

	config.AgentTimeout, err = strconv.Atoi(getEnv("AGENT_TIMEOUT", "600"))
	if err != nil {
		return nil, errors.New("wrong value for AGENT_TIMEOUT")
	}
	config.AgentPing, err = strconv.Atoi(getEnv("AGENT_PING", "60"))
	if err != nil {
		return nil, errors.New("wrong value for AGENT_PING")
	}
	config.AgentResolveTime, err = strconv.Atoi(getEnv("AGENT_RESOLVE_TIME", "600"))
	if err != nil {
		return nil, errors.New("wrong value for AGENT_RESOLVE_TIME")
	}

	return config, nil
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
