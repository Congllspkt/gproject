package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type UserTest1M struct {
	gorm.Model
	ID uint
	Name string
}

//1M -> 17s
func TryInsert1M() {
	dns := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True"
	db, _ := gorm.Open(mysql.Open(dns), &gorm.Config{})
	db.Migrator().HasTable(&UserTest1M{}) 
	db.AutoMigrate(&UserTest1M{})
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(10)
	defer sqlDb.Close()

	// Insert 1,000,000 records in batches
	var users []UserTest1M
	for i := 0; i < 1000000; i++ {
		users = append(users, UserTest1M{
			Name: fmt.Sprintf("User%d", i),
		})

		if len(users) >= 1000 { // Insert in batches of 1000
			db.CreateInBatches(users, 1000)
			users = users[:0] // Clear the slice for the next batch
		}
	}

	// Insert any remaining records
	if len(users) > 0 {
		db.CreateInBatches(users, 1000)
	}
}