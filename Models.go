package main

type Option struct {
	OptionId    string `json:"optionid,omitempty"`
	OptionsText string `json:"optiontext" bson:"optiontext"`
}

type Question struct {
	QuestionID   string   `json:"questionid,omitempty"`
	QuestionText string   `json:"questiontext"`
	Options      []Option `json:"options"`
	Answer       Option   `json:"answer"`
	Marks        int      `json:"marks"`
}

type Exam struct {
	ExamId    string   `json:"examid,omitempty"`
	ExamName  string   `json:"examname"`
	Questions []string `json:"questions"`
}

type Course struct {
	CourseID   string   `json:"courseid,omitempty"`
	CourseName string   `json:"coursename"`
	Exams      []string `json:"exams"`
}

/***************************************************************************************************/

type QuestionReport struct {
	QuestionID  string
	GivenAnswer string
	Marks       int
}

type ExamReport struct {
	ExamID          string
	QuestionReports []QuestionReport
}

type CourseReport struct {
	CourseID    string
	ExamReports []ExamReport
}

/***************************************************************************************************/

type Student struct {
	StudentID    string         `json:"studentid,omitempty"`
	StudentName  string         `json:"name"`
	College      string         `json:"college"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Course       []string       `json:"course"`
	CourseReport []CourseReport `json:"reports"`
}
