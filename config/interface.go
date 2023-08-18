package config

type envKey string

type Config struct {
	Port     string
	Database Database
	JWT      JWT
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWT struct {
	ExpireTime int
	SecretKey  string
	Cost       int
}
