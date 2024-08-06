package initialize

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserTest struct {
	gorm.Model
	ID   uint
	Name string
}

//1:10s
func TryInsert10M() {
	dns := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True"
	db, _ := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if db.Migrator().HasTable(&UserTest{}) {
		db.Migrator().DropTable(&UserTest{})
	}
	db.AutoMigrate(&UserTest{})
	sqlDb, _ := db.DB()
	numWorkers := 10 // Number of Goroutines to use
	sqlDb.SetMaxOpenConns(numWorkers)
	defer sqlDb.Close()

	// Insert 10,000,000 records using Goroutines
	const totalRecords = 10000000
	const batchSize = 1000

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	recordsPerWorker := totalRecords / numWorkers

	for w := 0; w < numWorkers; w++ {
		go func(workerID int) {
			defer wg.Done()

			var users []UserTest
			for i := workerID * recordsPerWorker; i < (workerID+1)*recordsPerWorker; i++ {
				users = append(users, UserTest{
					Name: fmt.Sprintf("ZUser%d", i),
				})

				if len(users) >= batchSize {
					db.CreateInBatches(users, batchSize)
					users = users[:0] // Clear the slice for the next batch
				}
			}

			// Insert any remaining records
			if len(users) > 0 {
				db.CreateInBatches(users, batchSize)

			}
		}(w)
	}
	wg.Wait()
}
