package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := retryConnectDB(dsn, 5, 2*time.Second)
	if err != nil {
		log.Fatal("Gagal connect ke DB setelah retry:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	fmt.Println("Connected to databasee............")
}


func retryConnectDB(dsn string, maxRetries int, wait time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		log.Println("Gagal connect ke DB, retry dalam", wait)
		time.Sleep(wait)
	}
	return nil, err
}
