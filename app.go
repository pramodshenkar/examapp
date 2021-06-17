package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/controllers"
	"github.com/pramodshenkar/examapp/middlewares"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/student")
	{
		v1.POST("/signup", controllers.Signup)
		v1.POST("/login", controllers.Login)
		v1.POST("/courses", controllers.GetSudentEnrolledCourses)
	}

	router.GET("/courses", controllers.GetAllCourses)
	router.POST("/course", controllers.GetCoursesByID)

	router.POST("/exams", controllers.GetExamsByCourseID)
	router.POST("/exam", controllers.GetExamsByExamID)

	router.POST("/report", controllers.GetReport)

	router.POST("/questions", controllers.GetQuestionsIDsByExamID)
	router.POST("/question", controllers.GetQuestionsByQuestionID)

	router.POST("/endexam", controllers.UpdateReportForEndExam)
	router.POST("/submitanswer", controllers.UpdateReportForSubmitAnswer)

	v2 := router.Group("")
	v2.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v2.Use(middlewares.Authenticate())
	{
		v2.POST("/hello", controllers.SayHello)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":5000")

}
