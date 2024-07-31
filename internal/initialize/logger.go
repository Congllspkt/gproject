package initialize

import (
	"gproject/internal/initialize/global"
	"gproject/internal/initialize/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger, "./storage/logs/dev_001.log")
	global.LoggerConsumer = logger.NewLogger(global.Config.Logger, "./storage/logs/Consumer_Received.log")
	global.LoggerProducer = logger.NewLogger(global.Config.Logger, "./storage/logs/Producer_Sent.log")
	global.Logger.Info("Init Logger success")
}
