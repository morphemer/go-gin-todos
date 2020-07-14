package models

import "github.com/jinzhu/gorm"

// User Userm model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
}
