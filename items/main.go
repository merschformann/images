package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes to handlers
	r.POST("/items", AddItem)
	r.GET("/items/:id", GetItem)
	r.GET("/items", ListItems)

	r.Run() // default listens on :8080
}
