package main

import (
	"go.uber.org/zap"
	"horizon-emulator/core/config"
	"horizon-emulator/core/logger"
)

func main() {

	tLog := logger.CreateTempLogger()
	tLog.Info("Starting Horizon emulator, please wait...")

	err := config.CreateDefaultConfig("config.ini", tLog)
	if err != nil {
		panic(err)
	}

	cfg, err := config.LoadConfig("config.ini", tLog)
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(cfg)
	zap.L().Info("Loaded configuration successfully")

}
