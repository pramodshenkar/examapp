package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetQuestionsByExamID(c *gin.Context) {
	// var data struct {
	// 	ExamID string `json:"exanmid" binding:"required"`
	// }

	// if c.BindJSON(&data) != nil {
	// 	fmt.Println("Provide required details")
	// 	c.JSON(400, gin.H{"message": "Provide required details"})
	// 	c.Abort()
	// 	return
	// }

	// questions, err := api.GetQuestionsByExamID(data.ExamID)

	// if err != nil {
	// 	fmt.Println("Problem to get exams ")
	// 	c.JSON(400, gin.H{"message": "Problem to get Exams"})
	// 	c.Abort()
	// 	return
	// }

	c.JSON(200, gin.H{"exam": "YET TO IMPLIMENT"})
}

func GetQuestionsByQuestionID(c *gin.Context) {
	var data struct {
		CourseId   string `json:"courseid" binding:"required"`
		QuestionID string `json:"questionid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	question, err := api.GetQuestionsByQuestionID(data.CourseId, data.QuestionID)

	if err != nil {
		fmt.Println("Problem to get exams ")
		c.JSON(400, gin.H{"message": "Problem to get Exams"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"exam": question})
}
