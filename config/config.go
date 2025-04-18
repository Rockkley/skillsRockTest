package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

type Config struct {
	DbConfig
	AppConfig
}

type AppConfig struct {
	AppName string
	AppPort string
}
type DbConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	return &Config{
		DbConfig{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "postgres"),
			DBName:     getEnv("DB_NAME", "taskdb"),
		},
		AppConfig{
			AppName: getEnv("APP_NAME", "TaskAppSkillsRock"),
			AppPort: getEnv("APP_PORT", "3000"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (c *Config) GetDBConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
