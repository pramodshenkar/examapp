package models

type QuestionReport struct {
	QuestionID  string `json:"optionid"`
	IsAnswered  bool   `json:"isanswered"`
	GivenAnswer string `json:"givenanswer"`
	Marks       int    `json:"marks"`
}

type ExamReport struct {
	ExamID          string           `json:"examid"`
	IsSubmitted     bool             `json:"issubmitted"`
	QuestionReports []QuestionReport `json:"questionreport"`
}

type CourseReport struct {
	CourseID    string       `json:"courseid"`
	ExamReports []ExamReport `json:"examreports"`
}

type StudentReport struct {
	StudentID string         `json:"studentid"`
	Report    []CourseReport `json:"report"`
}
