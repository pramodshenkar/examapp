package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetExamsByCourseID(c *gin.Context) {
	var data struct {
		CourseID string `json:"courseID" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	exams, err := api.GetExams(data.CourseID)

	if err != nil {
		fmt.Println("Problem to get exams ")
		c.JSON(400, gin.H{"message": "Problem to get Exams"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"exam": exams})
}

func GetExamsByExamID(c *gin.Context) {
	var data struct {
		CourseID string `json:"courseid" binding:"required"`
		ExamID   string `json:"examid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	exam, err := api.GetExamsByID(data.CourseID, data.ExamID)

	if err != nil {
		fmt.Println("Problem to get exams ")
		c.JSON(400, gin.H{"message": "Problem to get Exams"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"exam": exam})

}
