package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName string             `json:"coursename" bson:"coursename" binding:"required"`
	Exams      []Exam             `json:"exams" bson:"exams"`
}

type Exam struct {
	ExamID    string     `json:"examid" bson:"examid"`
	ExamName  string     `json:"examname" bson:"examname"`
	Attempts  int        `json:"attempts" bson:"attempts"`
	Questions []Question `json:"questions" bson:"questions"`
}

type Question struct {
	QuestionID   string   `json:"questionid" bson:"questionid"`
	QuestionType string   `json:"questiontype" bson:"questiontype"`
	QuestionText string   `json:"questiontext" bson:"questiontext"`
	Mediapath    string   `json:"mediapath" bson:"mediapath"`
	Filetype     string   `json:"filetype" bson:"filetype"`
	Options      []Option `json:"options" bson:"options"`
	Answer       []Option `json:"answer" bson:"answer"`
	Marks        int      `json:"marks" bson:"marks"`
}

type Option struct {
	OptionId    string `json:"optionid" bson:"optionid"`
	OptionsText string `json:"optiontext" bson:"optiontext"`
}
