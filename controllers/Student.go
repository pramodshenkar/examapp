package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
	"github.com/pramodshenkar/examapp/models"
)

func Signup(c *gin.Context) {

	var student models.Student

	if c.BindJSON(&student) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	_, err := api.AddStudent(student)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Record Added Successfully"})
}

func Login(c *gin.Context) {
	c.JSON(201, gin.H{"message": "NOT IMPLIMENTED YET"})
}
