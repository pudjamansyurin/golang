package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDatabase(dbname string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbname)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Product{})

	return db
}
