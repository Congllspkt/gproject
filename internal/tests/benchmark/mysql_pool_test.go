package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Tipjs "}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}


func BenchmarkMaxOpenConns1(b *testing.B) {
	benchmarkMaxOpenConns(b,1)
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	benchmarkMaxOpenConns(b,10)
}




func benchmarkMaxOpenConns(b *testing.B, maxOpen int) {
	dns := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{

	})

	if err != nil {
		log.Fatalf("fail connect db %v", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			log.Fatalf("fail drop table %v", err)
		}
	}

	db.AutoMigrate(&User{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail get DB %v", err)
	}

	sqlDb.SetMaxOpenConns(maxOpen)
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

