package models

import "github.com/jinzhu/gorm"

// Article article
type Article struct {
	gorm.Model
	UserID  int
	User    User
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}
