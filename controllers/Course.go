package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func ShowCourses(c *gin.Context) {
	courses, err := api.GetCourses()

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"courses": courses})

}
