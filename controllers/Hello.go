package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	var student struct {
		Name string `json:"name"`
	}

	if c.BindJSON(&student) != nil {
		c.JSON(400, gin.H{"message": "Provide required details."})
		c.Abort()
		return
	}

	message := fmt.Sprintf("%s%s", "Good Morning ", student.Name)

	c.JSON(200, gin.H{"message": message})
}
