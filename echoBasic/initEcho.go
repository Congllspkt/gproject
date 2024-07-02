package echobasic

import (
	"gproject/db"
	"gproject/handler"
	"gproject/repository/repo_impl"

	"github.com/labstack/echo/v4"
)

func InitEcho() {
	sql := &db.Sql{
		Host: "localhost",
		Port:5432,
		UserName:"myuser",
		Password:"mypassword",
		DbName:"mydb",
	}
	sql.Connect()
	defer sql.Close();



	e := echo.New()
	userHandler = handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	e.Logger.Fatal(e.Start(":3000"))
}