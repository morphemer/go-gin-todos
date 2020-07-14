package models

// User Userm model
type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
}
