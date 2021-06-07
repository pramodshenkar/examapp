package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/hello", controllers.Hello)

	v1 := router.Group("")
	{
		v1.POST("/signup", controllers.Signup)
		v1.POST("/login", controllers.Login)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":5000")
}
