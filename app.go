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

	v1 := router.Group("/student")
	{
		v1.POST("/signup", controllers.Signup)
		v1.POST("/login", controllers.Login)
		v1.POST("/courses", controllers.GetSudentEnrolledCourses)

	}
	router.GET("/courses", controllers.GetAllCourses)

	router.POST("/exams", controllers.GetExamsByCourseID)
	router.POST("/exam", controllers.GetExamsByExamID)
	router.POST("/questions", controllers.GetQuestionsByExamID)
	router.POST("/question", controllers.GetQuestionsByQuestionID)

	router.POST("/report", controllers.GetStudentReport)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":5000")

}
