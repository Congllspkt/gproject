package initialize

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserTest struct {
	gorm.Model
	ID uint
	Name string
}

//1:10s
func TryInsert10M() {
	dns := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("fail connect db %v", err)
	}

	if db.Migrator().HasTable(&UserTest{}) {
		if err := db.Migrator().DropTable(&UserTest{}); err != nil {
			log.Fatalf("fail drop table %v", err)
		}
	}

	db.AutoMigrate(&UserTest{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail get DB %v", err)
	}

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
					// Add other fields as needed
				})

				if len(users) >= batchSize {
					result := db.CreateInBatches(users, batchSize)
					if result.Error != nil {
						log.Fatalf("fail to insert records in batch: %v", result.Error)
					}
					users = users[:0] // Clear the slice for the next batch
				}
			}

			// Insert any remaining records
			if len(users) > 0 {
				result := db.CreateInBatches(users, batchSize)
				if result.Error != nil {
					log.Fatalf("fail to insert remaining records in batch: %v", result.Error)
				}
			}
		}(w)
	}

	wg.Wait()
}