package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MyHandler is the interface
type MyHandler interface {
	HelloWorld(c *gin.Context)
}

type myHandler struct {
}

// NewMyHandler is a function new instance
func NewMyHandler() MyHandler {
	return &myHandler{}
}

func (h *myHandler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Hello World! %s", time.Now()),
	})
}
