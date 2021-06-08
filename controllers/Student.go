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
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	result, err := api.GetStudentByUsername(student.Username)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem While Cheking username"})
		c.Abort()
		return
	}

	if result.Username != "" {
		c.JSON(409, gin.H{"message": "Account Already exist"})
		c.Abort()
		return
	}

	studentid, err := api.AddStudent(student)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}
	fmt.Println(studentid, "added to database")
	c.JSON(200, gin.H{"message": "Record Added Successfully"})
}

func Login(c *gin.Context) {
	var data models.StudentCredentials

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := api.GetStudentByUsername(data.Username)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	if result.Username == "" {
		c.JSON(404, gin.H{"message": "Opps! Username is not found"})
		c.Abort()
		return
	}

	fmt.Println(result)

	// hashedPassword := []byte(result.Password)
	// password := []byte(data.Password)

	// err = helpers.PasswordCompare(password, hashedPassword)

	// if err != nil {
	// 	c.JSON(403, gin.H{"message": "Invalid user credentials"})
	// 	c.Abort()
	// 	return
	// }

	// jwtToken, err2 := services.GenerateToken(data.Email)

	// // If we fail to generate token for access
	// if err2 != nil {
	// 	c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
	// 	c.Abort()
	// 	return
	// }

	if result.Password != data.Password {
		c.JSON(404, gin.H{"message": "Opps Wrong password"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": result})
}
