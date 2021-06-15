package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
	"github.com/pramodshenkar/examapp/models"
)

func Signup(c *gin.Context) {

	var student models.Student

	if c.BindJSON(&student) != nil {
		c.JSON(400, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	credentials, err := api.GetStudentCredentials(student.Username)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem While Cheking if username exists"})
		c.Abort()
		return
	}

	if credentials.Username != "" {
		c.JSON(409, gin.H{"message": "Account Already exist"})
		c.Abort()
		return
	}

	studentid := api.GenerateStudentID()

	isSavedData := api.AddStudent(studentid, student)

	isSavedCredentials := api.AddCredentials(studentid, student)

	if !isSavedData || !isSavedCredentials {
		c.JSON(400, gin.H{"message": "Cannot save Credentials"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Data saved successfully"})

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem to write credentials."})
		c.Abort()
		return
	}
}

func Login(c *gin.Context) {
	var getStudent struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.BindJSON(&getStudent) != nil {
		fmt.Println("Provide required details")
		c.JSON(400, gin.H{"message": "Provide required details."})
		c.Abort()
		return
	}

	studentCredentials, err := api.GetStudentCredentials(getStudent.Username)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem to get credentials."})
		c.Abort()
		return
	}

	if studentCredentials.Username == "" {
		fmt.Println("Opps! Username is not found.")
		c.JSON(404, gin.H{"message": "Opps! Username is not found."})
		c.Abort()
		return
	}

	if studentCredentials.Password != getStudent.Password {
		c.JSON(404, gin.H{"message": "Opps! Wrong password."})
		c.Abort()
		return
	}

	student, err := api.GetStudent(studentCredentials.Username)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem to your data."})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"student": student})
}
