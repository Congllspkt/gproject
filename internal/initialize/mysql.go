package initialize

import (
	"fmt"
	"gproject/internal/initialize/global"
	"gproject/internal/initialize/po"
	"gproject/internal/initialize/setting"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
	}
}

func InitMySql() {
	m := global.Config.MySQL
	dns := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dns, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "Init mySQL fail")
	global.Logger.Info("Init mySQL success")
	global.Mdb = db

	setPool(m)
	migrateTables()
}

func setPool(m setting.MySQLSetting) {
	sqlDb, err := global.Mdb.DB()

	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConMaxLifeTime))

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("Migrating tables err: ", err)
	}

}