package config

const (
	PORT envKey = "PORT"

	DB_HOST     envKey = "DB_HOST"
	DB_PORT     envKey = "DB_PORT"
	DB_USER     envKey = "DB_USER"
	DB_PASSWORD envKey = "DB_PASSWORD"
	DB_NAME     envKey = "DB_NAME"
	DB_SSL_MODE envKey = "DB_SSL_MODE"

	JWT_EXPIRE_TIME envKey = "JWT_EXPIRE_TIME"
	JWT_SECRET_KEY  envKey = "JWT_SECRET_KEY"
	JWT_COST        envKey = "JWT_COST"
)
