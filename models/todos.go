package models

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joe41203/go-gin-todos/config"
)

// GetAllTodos fetch all of todos
func GetAllTodos(todos *[]Todo) (err error) {
	if err = config.DB.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

// GetATodo find a todo
func GetATodo(todo *Todo, id string) (err error) {
	var intID int
	intID, _ = strconv.Atoi(id)
	if err = config.DB.First(&todo, intID).Error; err != nil {
		return err
	}

	return nil
}
