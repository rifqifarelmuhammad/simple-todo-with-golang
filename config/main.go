package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

var config = &Config{}

func GetInstance() *Config {
	return config
}

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(constant.TAG_CONFIG, err, "Failed to load .env file")
	}

	config = GenerateConfig()
	log.Print(constant.TAG_CONFIG, "Load Environment Variables Completed")
}

func GenerateConfig() *Config {
	return &Config{
		Server: Server{
			Port:      GetStringEnv(APP_PORT),
			Whitelist: GetArrayOfStringEnv(APP_WHITELIST),
		},
		Database: Database{
			Host:     GetStringEnv(DB_HOST),
			Port:     StringToInt(GetStringEnv(DB_PORT), 5432),
			User:     GetStringEnv(DB_USER),
			Password: GetStringEnv(DB_PASSWORD),
			Name:     GetStringEnv(DB_NAME),
			SSLMode:  GetStringEnv(DB_SSL_MODE),
		},
		JWT: JWT{
			ExpireTime: StringToInt(GetStringEnv(JWT_EXPIRE_TIME), 7),
			SecretKey:  GetStringEnv(JWT_SECRET_KEY),
			Cost:       StringToInt(GetStringEnv(JWT_COST), 7),
		},
	}
}

func StringToInt(str string, alternate int) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return alternate
	}
	return result
}

func GetStringEnv(key envKey) string {
	return os.Getenv(string(key))
}

func GetArrayOfStringEnv(key envKey) []string {
	return strings.Split(os.Getenv(string(key)), ",")
}
