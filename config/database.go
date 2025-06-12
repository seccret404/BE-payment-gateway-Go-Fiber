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

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Tidak dapat menemuan database url")
	}

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Gagal connect ke DB")
	}

	sqlDB, err := db.DB()
	if err != nil{
		log.Fatal("error", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	fmt.Println("Connected to databasee............")
}