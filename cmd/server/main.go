package main

import (
	// "gproject/internal/cli"
	"gproject/internal/cli/viper_conf"
	// "gproject/internal/routers"
)

func main() {

	// r := routers.NewRouter()
	// r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// cli.LogMain()
	viper_conf.GetViperConf()
}
