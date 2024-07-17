package initialize

import (
	"gproject/internal/initialize/global"
	"gproject/internal/initialize/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger()
}
