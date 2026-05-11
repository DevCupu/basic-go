package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
