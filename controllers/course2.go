package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
	"github.com/pramodshenkar/examapp/models"
)

func GetAllCourses(c *gin.Context) {
	categories, err := api.GetAllCourses()

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"courses": categories})
}

func AddCourse(c *gin.Context) {

	course := models.Course{}

	err := c.BindJSON(&course)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for AddCourse"})
		c.Abort()
		return
	}

	result, err := api.AddCourse(course)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"course": result})

}
