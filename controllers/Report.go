package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
)

// func GenerateReport(c *gin.Context) {

// 	var data struct {
// 		Username string `json:"username" binding:"required"`
// 	}

// 	if c.BindJSON(&data) != nil {
// 		fmt.Println("Provide required details")
// 		c.JSON(400, gin.H{"message": "Provide required details"})
// 		c.Abort()
// 		return
// 	}

// 	student, err := api.GetStudentByUsername(data.Username)

// 	if err != nil {
// 		fmt.Println("Problem logging into your account")
// 		c.JSON(400, gin.H{"message": "Problem logging into your account"})
// 		c.Abort()
// 		return
// 	}

// 	if student.Username == "" {
// 		fmt.Println("Opps! Username is not found")

// 		c.JSON(404, gin.H{"message": "Opps! Username is not found"})
// 		c.Abort()
// 		return
// 	}

// 	fmt.Println(student)

// 	courseReports, _ := api.GenerateCourseReport(student.Courses)

// 	c.JSON(200, gin.H{"message": courseReports})
// }

func GetStudentReport(c *gin.Context) {

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
		fmt.Println("Problem lwhile getting username")
		c.JSON(400, gin.H{"message": "Problem while getting username"})
		c.Abort()
		return
	}

	report, err := api.GetStudentReport(student.StudentID)

	if err != nil {
		fmt.Println("Problem while Getting report")
		c.JSON(400, gin.H{"message": "Problem while Getting report"})
		c.Abort()
		return
	}

	// if report.StudentID == "" {
	// 	fmt.Println("Opps! Report id is not matched")

	// 	c.JSON(404, gin.H{"message": "Opps! Username is not found"})
	// 	c.Abort()
	// 	return
	// }

	fmt.Println(report)

	c.JSON(200, gin.H{"report": report.Report})
}
