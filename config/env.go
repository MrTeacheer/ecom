package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST"),
		Port:       getEnv("PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:     getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return "not taken"
}
