package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// User Userm model
type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null" validate:"required"`
	Email    string `gorm:"column:email;unique_index;not null" validate:"required,email"`
}

var validate *validator.Validate

// Validate user validation
func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
