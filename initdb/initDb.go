package initdb

import "gproject/db"

func InitDB() {
	sql := &db.Sql{
		Host: "localhost",
		Port:5432,
		UserName:"myuser",
		Password:"mypassword",
		DbName:"mydb",
	}
	sql.Connect()
	defer sql.Close();
}