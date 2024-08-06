package initialize

import (
	"fmt"
	"log"

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
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("fail connect db %v", err)
	}

	if db.Migrator().HasTable(&UserTest1M{}) {
		if err := db.Migrator().DropTable(&UserTest1M{}); err != nil {
			log.Fatalf("fail drop table %v", err)
		}
	}

	db.AutoMigrate(&UserTest1M{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail get DB %v", err)
	}

	sqlDb.SetMaxOpenConns(10)
	defer sqlDb.Close()

	// Insert 1,000,000 records in batches
	var users []UserTest1M
	for i := 0; i < 1000000; i++ {
		users = append(users, UserTest1M{
			Name: fmt.Sprintf("User%d", i),
			// Add other fields as needed
		})

		if len(users) >= 1000 { // Insert in batches of 1000
			result := db.CreateInBatches(users, 1000)
			if result.Error != nil {
				log.Fatalf("fail to insert records in batch: %v", result.Error)
			}
			users = users[:0] // Clear the slice for the next batch
		}
	}

	// Insert any remaining records
	if len(users) > 0 {
		result := db.CreateInBatches(users, 1000)
		if result.Error != nil {
			log.Fatalf("fail to insert remaining records in batch: %v", result.Error)
		}
	}
}