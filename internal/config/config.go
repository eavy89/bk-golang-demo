package config

import (
	"fmt"
	"os"
)

const (
	ClaimUserID   = "user_id"
	ClaimUsername = "username"

	EnvFile = "/app/.env"

	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	DB_PATH        = "DB_PATH"
	SERVER_PORT    = "SERVER_PORT"
	POD_NAME       = "POD_NAME"
)

func GetJWTKey() ([]byte, error) {
	//envFile, _ := godotenv.Read(EnvFile) // we can load this file in memory
	//value := envFile["JWT_SECRET_KEY"]

	value := os.Getenv(JWT_SECRET_KEY)
	if value == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	return []byte(value), nil
}

func GetDBPath() string {
	//envFile, _ := godotenv.Read(EnvFile) // we can load this file in memory
	//value := envFile["DB_PATH"]

	value := os.Getenv(DB_PATH)
	return value
}

func GetServerPort() string {
	value := os.Getenv(SERVER_PORT)
	if value == "" {
		value = ":8080"
	}
	return ":" + value
}

func GetLogFilename() string {
	value := os.Getenv(POD_NAME)
	if value == "" {
		value = "00"
	}
	return "data/log-" + value + ".log"
}
