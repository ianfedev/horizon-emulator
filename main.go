package main

import (
	"go.uber.org/zap"
	"horizon-emulator/core/config"
	"horizon-emulator/core/logger"
)

func main() {

	tLog := logger.CreateTempLogger()
	tLog.Info("Starting Horizon emulator, please wait...")

	cfg, err := config.LoadConfig("config.ini", tLog)
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(cfg)
	zap.L().Error("XD")

}
