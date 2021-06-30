package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Learner struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	LearnerName string             `json:"learnername" bson:"learnername" binding:"required"`
	Username    string             `json:"username" bson:"username" binding:"required"`
	Password    string             `json:"password" bson:"password" binding:"required"`
	College     string             `json:"college" bson:"college"`
	Email       string             `json:"email" bson:"email" binding:"required,email"`
	Courses     []Course           `json:"courses" bson:"courses" binding:"required"`
	Report      []ExamReport       `json:"examreport" bson:"examreport"`
}
