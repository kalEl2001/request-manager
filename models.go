package main

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Request struct {
	gorm.Model
	User string
	Slug string
	Status uint
	Files []FileQueue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type FileQueue struct {
	gorm.Model
	RequestId uint
	link string
}

func MigrateDB() {
	DB_URL := os.Getenv("DATABASE_URL")
	if len(DB_URL) == 0 {
		DB_URL = "data.db"
	}
	db, err := gorm.Open(sqlite.Open(DB_URL), &gorm.Config{})
	if err != nil {
		field := map[string]interface{}{
			"db_url": DB_URL,
			"error": err,
		}
		createLog(nil, "Panic", "Failed to connect to database", field)
	}
	createLog(nil, "Info", "Successfully to connect to database", nil)

	// Migrate the schema
	db.AutoMigrate(&Request{})
	db.AutoMigrate(&FileQueue{})
}