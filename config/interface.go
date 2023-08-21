package config

type envKey string

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
}

type Server struct {
	Port      string
	Whitelist []string
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
