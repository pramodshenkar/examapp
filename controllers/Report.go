package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetReport(c *gin.Context) {

	var data struct {
		StudentID string `json:"studentid" binding:"required"`
		CourseID  string `json:"courseid" binding:"required"`
		ExamID    string `json:"examid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	examReport, err := api.GetReport(data.StudentID, data.CourseID, data.ExamID, 3)

	if err != nil {
		fmt.Println("Problem to get examreport file")
		c.JSON(400, gin.H{"message": "Problem to get examreport file"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"examreport": examReport})

}

func UpdateReportForEndExam(c *gin.Context) {

	var data struct {
		StudentID string `json:"studentid" binding:"required"`
		CourseID  string `json:"courseid" binding:"required"`
		ExamID    string `json:"examid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	isUpdated := api.UpdateReportForEndExam(data.StudentID, data.CourseID, data.ExamID)

	if !isUpdated {
		c.JSON(400, gin.H{"message": "Cant update report"})
		c.Abort()
		return

	}

	c.JSON(200, gin.H{"EndExam": isUpdated})

}

func UpdateReportForSubmitAnswer(c *gin.Context) {

	var data struct {
		StudentID  string `json:"studentid" binding:"required"`
		CourseID   string `json:"courseid" binding:"required"`
		ExamID     string `json:"examid" binding:"required"`
		QuestionID string `json:"questionid" binding:"required"`
		AnswerID   string `json:"answerid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	isUpdated := api.UpdateReportForSubmitAnswer(data.StudentID, data.CourseID, data.ExamID, data.QuestionID, data.AnswerID)

	if !isUpdated {
		c.JSON(400, gin.H{"message": "Cant update report"})
		c.Abort()
		return

	}

	c.JSON(200, gin.H{"SubmitAnswer": isUpdated})

}
