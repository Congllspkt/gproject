package initialize

import (
	"fmt"
	"gproject/internal/initialize/global"

	"go.uber.org/zap"
)

func Run() {

	InitConFig()
	m := global.Config.MySQL
	fmt.Println("-----------", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))

	InitMySql()
	InitRedis()
	InitKafka()

	r := InitRouter()

	r.Run(":8002")

}