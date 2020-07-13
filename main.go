package main

import (
	"fmt"

	"github.com/joe41203/go-gin-todos/config"
)

func main() {
	fmt.Println("HelloWorld")
	db := config.Init()
	defer db.Close()
}
