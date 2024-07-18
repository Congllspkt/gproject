package global

import (
	"gproject/internal/initialize/logger"
	"gproject/internal/initialize/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)