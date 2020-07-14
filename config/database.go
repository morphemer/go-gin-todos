package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Database type
type Database struct {
	*gorm.DB
}

// DB gorm.DB
var DB *gorm.DB

// Init return DB
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gorm.db")

	if err != nil {
		fmt.Println("DB Error", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)

	DB = db
	return DB
}

// GetDB get DB
func GetDB() *gorm.DB {
	return DB
}
