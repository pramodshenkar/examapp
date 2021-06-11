package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetAllCourses(c *gin.Context) {
	courses, err := api.GetCourses()

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem to get Courses"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"courses": courses})
}

func GetSudentEnrolledCourses(c *gin.Context) {
	var data struct {
		Username string `json:"username" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	student, err := api.GetStudent(data.Username)

	if err != nil {
		fmt.Println("Problem logging into your account")
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	if student.Username == "" {
		fmt.Println("Opps! Username is not found")

		c.JSON(404, gin.H{"message": "Opps! Username is not found"})
		c.Abort()
		return
	}

	courses, err := api.GetCoursesByUsername(student)

	if err != nil {
		fmt.Println("Problem While getting courses")
		c.JSON(400, gin.H{"message": "Problem While getting courses"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"courses": courses})
}
