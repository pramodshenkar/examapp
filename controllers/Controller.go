package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	fmt.Println("hello user")
}
