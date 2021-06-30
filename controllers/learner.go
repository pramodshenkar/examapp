package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/examapp/api"
	"github.com/pramodshenkar/examapp/models"
	"github.com/pramodshenkar/examapp/services"
)

func AddLearner(c *gin.Context) {

	learner := models.Learner{}

	err := c.BindJSON(&learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for AddLearner"})
		c.Abort()
		return
	}

	result, err := api.AddLearner(learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"learner": result})

}

func GetAllLearners(c *gin.Context) {
	learners, err := api.GetAllLearners()

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for GetAllLearner"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"learners": learners})
}

func GetLearner(c *gin.Context) {

	var learnerRequest struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.BindJSON(&learnerRequest); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for GetLearner"})
		c.Abort()
		return
	}

	learner, err := api.GetLearner(learnerRequest.Username)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"learner": learner})

}

func DeleteLearner(c *gin.Context) {
	var learner struct {
		LearnerID string `json:"_id" bson:"_id"`
	}

	if err := c.BindJSON(&learner); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for DeleteLearner"})
		c.Abort()
		return
	}

	result, err := api.DeleteLearner(learner.LearnerID)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Sorry. we can't able to Delete Learner Information"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"learner": result})
}

func UpdateLearner(c *gin.Context) {

	learner := models.Learner{}

	err := c.BindJSON(&learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for UpdateLearner"})
		c.Abort()
		return
	}

	result, err := api.UpdateLearner(learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Sorry. we can't able to update Learner Information"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"learner": result})

}

func Signup(c *gin.Context) {

	learner := models.Learner{}

	err := c.BindJSON(&learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details for Signup User"})
		c.Abort()
		return
	}

	existLearner, err := api.GetLearner(learner.Username)

	if existLearner.Username != "" {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Accout Already present"})
		c.Abort()
		return
	}

	result, err := api.AddLearner(learner)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"learner": result})
}

func Login(c *gin.Context) {
	var learnerRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.BindJSON(&learnerRequest) != nil {
		fmt.Println("Provide required details for Login")
		c.JSON(400, gin.H{"message": "Provide required details."})
		c.Abort()
		return
	}

	existLearner, err := api.GetLearner(learnerRequest.Username)

	if existLearner.Username == "" {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Account not found"})
		c.Abort()
		return
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Problem to get your data."})
		c.Abort()
		return
	}

	if learnerRequest.Username != existLearner.Username {
		fmt.Println("Opps! Username is not found.")
		c.JSON(404, gin.H{"message": "Opps! Username is not found."})
		c.Abort()
		return
	}

	if learnerRequest.Password != existLearner.Password {
		c.JSON(404, gin.H{"message": "Opps! Wrong password."})
		c.Abort()
		return
	}

	token, err := services.GenerateToken(existLearner.Username)

	if err != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"token": token})
}
