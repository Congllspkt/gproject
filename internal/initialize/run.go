package initialize

import (
)

func Run() {
	InitConFig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	tryMySQL()
}
