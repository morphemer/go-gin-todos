package models

import "github.com/jinzhu/gorm"

// Article article
type Article struct {
	gorm.Model
	UserID  int
	User    User
	Title   string `gorm:"column:title;not null"`
	Content string `gorm:"column:content;not null"`
}
