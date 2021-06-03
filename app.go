package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/controllers"
)

func main() {
	router := gin.Default()
	v1 := router.Group("")
	{
		student := new(controllers.StudentController)
		v1.POST("/signup", student.Signup)
		v1.POST("/login", student.Login)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":5000")
}
