package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joe41203/go-gin-todos/config"
)

// GetAllTodos fetch all of todos
func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
