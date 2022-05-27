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
	Status int
	NumFiles int
	Files []FileQueue `gorm:"foreignKey:RequestId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type FileQueue struct {
	gorm.Model
	RequestId uint
	Link string
	Status bool
}

var dbConn *gorm.DB

func initDBConnection() {
	DB_URL := os.Getenv("DATABASE_URL")
	if len(DB_URL) == 0 {
		DB_URL = "data.db"
	}
	
	db, err := gorm.Open(sqlite.Open(DB_URL), &gorm.Config{})
	if err != nil {
		failLog(err, "Failed to connect to database")
	}
	dbConn = db
	infoLog("Successfully connect to database", nil)
}

func migrateDB() {
	// Migrate the schema
	dbConn.AutoMigrate(&Request{})
	dbConn.AutoMigrate(&FileQueue{})
	infoLog("Successfully migrating database", nil)
}