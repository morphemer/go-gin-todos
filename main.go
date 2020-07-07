package main

import (
	routes "github.com/joe41203/go-gin-todos/config"
)

func main() {
	r := routes.Routes()
	r.Run()
}
