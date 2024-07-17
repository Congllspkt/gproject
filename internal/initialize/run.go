package initialize

import (
	"fmt"
	"gproject/internal/initialize/global"
)

func Run() {

	InitConFig()
	m := global.Config.MySQL
	fmt.Println("-----------", m.Username, m.Password)
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	r := InitRouter()

	r.Run(":8002")

}