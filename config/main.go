package config

import (
	"os"
	"strconv"

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
		Port: GetEnv(PORT),
		Database: Database{
			Host:     GetEnv(DB_HOST),
			Port:     StringToInt(GetEnv(DB_PORT), 5432),
			User:     GetEnv(DB_USER),
			Password: GetEnv(DB_PASSWORD),
			Name:     GetEnv(DB_NAME),
			SSLMode:  GetEnv(DB_SSL_MODE),
		},
		JWT: JWT{
			ExpireTime: StringToInt(GetEnv(JWT_EXPIRE_TIME), 7),
			SecretKey:  GetEnv(JWT_SECRET_KEY),
			Cost:       StringToInt(GetEnv(JWT_COST), 7),
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

func GetEnv(key envKey) string {
	return os.Getenv(string(key))
}
