package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB gorm
var DB *gorm.DB

// DBConfig db config
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig build config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "go_gin_todos",
		Password: "",
	}
	return &dbConfig
}

// DbURL db url
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
