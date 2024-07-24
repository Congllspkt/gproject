package initialize

func Run() {

	InitConFig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	TryDataSample()
	r := InitRouter()

	r.POST("action/stock", actionStock)

	go RegisterConsumeATC(1)
	go RegisterConsumeATC(2)

	r.Run(":8999")

}
