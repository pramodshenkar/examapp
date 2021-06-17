package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetQuestionsIDsByExamID(c *gin.Context) {
	var data struct {
		CourseID string `json:"courseid" binding:"required"`
		ExamID   string `json:"examid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details for GetQuestionsIDsByExamID")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	questions, err := api.GetQuestionsIDsByExamID(data.CourseID, data.ExamID)

	if err != nil {
		fmt.Println("Problem to get questions for intended examid ")
		c.JSON(400, gin.H{"message": "Problem to get Problem to get questions for intended examid"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"questions": questions})
}

func GetQuestionsByQuestionID(c *gin.Context) {
	var data struct {
		CourseId   string `json:"courseid" binding:"required"`
		QuestionID string `json:"questionid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details for GetQuestionsByQuestionID")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	question, err := api.GetQuestionsByQuestionID(data.CourseId, data.QuestionID)

	if err != nil {
		fmt.Println("Problem to get questions for intended questionid ")
		c.JSON(400, gin.H{"message": "Problem to get questions for intended questionid"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"question": question})
}
