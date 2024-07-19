package initialize

func Run() {

	InitConFig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	TryDataSample()
	// r := InitRouter()
	// r.Run(":8002")

}
