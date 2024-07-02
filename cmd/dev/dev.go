package cmd

import (
	"fmt"
	"gproject/db"
	_ "gproject/docs"
	"gproject/handler"
	"gproject/helper"
	"gproject/log"
	"gproject/repository/repo_impl"
	"gproject/router"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func init() {
	fmt.Println("DEV ENVIROMENT")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

// @title Github Trending API
// @version 1.0
// @description More
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey jwt
// @in header
// @name Authorization

// @host localhost:3000
// @BasePath /
func MainDev() {
	sql := &db.Sql{
		Host:     "localhost", //localhost,host.docker.internal
		Port:     5432,
		UserName: "myuser",
		Password: "mypassword",
		DbName:   "mydb",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()

	go scheduleUpdateTrending(360*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
