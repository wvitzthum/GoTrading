package main

import (
	"GoTrading/app"
	"GoTrading/pkg/config"
	"GoTrading/pkg/log"
	"flag"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

func main() {
	logger := logger.InitSlog(slog.LevelInfo)
	envErr := godotenv.Load()
	if envErr != nil {
		logger.Error("Error loading envs from .env file", nil)
		return
	}


	config, err := config.GetConfig("./config/" ,*flag.String("configname", "default", "Config file"))
	if err != nil {
		logger.Error("Could not read config", err)
		return
	}

	err = app.InitApp(config, logger)
	if err != nil {
		logger.Error("Failed to init Application", err)
	}
}