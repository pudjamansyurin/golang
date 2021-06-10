package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase(dbname string) {
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db.AutoMigrate(&Product{})
}
