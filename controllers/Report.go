package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

func GetReport(c *gin.Context) {

	var data struct {
		UserID   string `json:"userid" binding:"required"`
		CourseID string `json:"courseid" binding:"required"`
		ExamID   string `json:"examid" binding:"required"`
	}

	if c.BindJSON(&data) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	examReport, err := api.GetReport(data.UserID, data.CourseID, data.ExamID)

	if err != nil {
		fmt.Println("Problem to get examreport file")
		c.JSON(400, gin.H{"message": "Problem to get examreport file"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"examreport": examReport})

}
