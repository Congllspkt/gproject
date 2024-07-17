package global

import (
	"gproject/internal/initialize/logger"
	"gproject/internal/initialize/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)