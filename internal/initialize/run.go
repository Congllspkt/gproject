package initialize

func Run() {

	InitConFig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	r := InitRouter()

	r.Run(":8002")

}