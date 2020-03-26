package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	handler := NewMyHandler()
	router.GET("/", handler.HelloWorld)

	router.Run(fmt.Sprintf(":%s", port))
}
