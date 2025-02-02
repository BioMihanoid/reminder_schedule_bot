package config

import (
	"github.com/joho/godotenv"
	"log/slog"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error(err.Error())
	}
}
