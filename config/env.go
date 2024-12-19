package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTexparation int64
	JWTsecret string
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
		JWTexparation: getEnvAsInt("JWT_EXP"),
		JWTsecret: getEnv("JWTsecret"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return "not taken"
}

func getEnvAsInt(key string) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0
		}

		return i
	}

	return 0
}