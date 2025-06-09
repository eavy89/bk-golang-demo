package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

const (
	ClaimUserID   = "user_id"
	ClaimUsername = "username"

	EnvFile = "/app/.env"
)

func GetJWTKey() []byte {
	envFile, _ := godotenv.Read(EnvFile) // we can load this file in memory
	value := envFile["JWT_SECRET_KEY"]

	fmt.Println("JWT_SECRET_KEY: ", value)
	return []byte(value)
}

func GetDBPath() string {
	envFile, _ := godotenv.Read(EnvFile) // we can load this file in memory
	value := envFile["DB_PATH"]
	return value
}
