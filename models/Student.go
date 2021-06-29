package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Learner struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	LearnerName string             `json:"learnername" bson:"learnername" binding:"required"`
	Username    string             `json:"username" bson:"username" binding:"required"`
	Password    string             `json:"password" bson:"password" binding:"required"`
	College     string             `json:"college" bson:"college"`
	Email       string             `json:"email" bson:"email" binding:"required,email"`
	Courses     []string           `json:"courses" bson:"courses"`
	Report      []ExamReport       `json:"examreport" bson:"examreport"`
}

type Student struct {
	StudentID   string   `json:"studentid"`
	StudentName string   `json:"name" binding:"required"`
	Username    string   `json:"username" binding:"required"`
	College     string   `json:"college"`
	Email       string   `json:"email" binding:"required,email"`
	Password    string   `json:"password" binding:"required"`
	Courses     []string `json:"courses"`
}

type Credentials struct {
	StudentID string `json:"studentid"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
