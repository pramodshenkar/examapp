package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
	"github.com/pramodshenkar/examapp/services"
)

func Authenticate() gin.HandlerFunc {

	return func(c *gin.Context) {

		requiredToken := c.Request.Header["Authorization"]

		if len(requiredToken) == 0 {
			c.JSON(400, gin.H{"message": "Please login your account"})
			c.Abort()
			return
		}

		studentID, _ := services.DecodeToken(requiredToken[0])

		result, err := api.GetStudent(studentID)
		if result.Username == "" {
			c.JSON(400, gin.H{"message": "User account not found"})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(400, gin.H{"message": "Something went wrong"})
			c.Abort()
			return
		}
		c.Set("User", result)
		c.Next()
	}
}
