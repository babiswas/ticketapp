package helper

import (
	"tokenapp/logger"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	logger := logger.LoggingInit()
	logger.Info("loading the enviroment variables.")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load the enviroment variables.")
	}
}
