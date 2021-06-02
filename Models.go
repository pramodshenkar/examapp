package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Option struct {
	OptionId    bson.ObjectId `json:"optionid,omitempty"`
	OptionsText string        `json:"optiontext" bson:"optiontext"`
}

type Question struct {
	QuestionID   bson.ObjectId `json:"questionid,omitempty"`
	QuestionText string        `json:"questiontext"`
	Options      []Option      `json:"options"`
	Answer       Option        `json:"answer"`
	Marks        int           `json:"marks"`
}

type Exam struct {
	ExamId    bson.ObjectId   `json:"examid,omitempty"`
	ExamName  string          `json:"examname"`
	Questions []bson.ObjectId `json:"questions"`
}

type Course struct {
	CourseID   bson.ObjectId   `json:"courseid,omitempty"`
	CourseName string          `json:"coursename"`
	Exams      []bson.ObjectId `json:"exams"`
}

type Student struct {
	StudentID   bson.ObjectId   `json:"studentid,omitempty"`
	StudentName string          `json:"name"`
	College     string          `json:"college"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	Course      []bson.ObjectId `json:"course"`
}
