package models

type QuestionReport struct {
	QuestionID  string `json:"optionid"`
	IsAnswered  bool   `json:"isanswered"`
	GivenAnswer string `json:"givenanswer"`
	Marks       int    `json:"marks"`
}

type AttemptReport struct {
	AttemptNo      int              `json:"attemptno"`
	IsSubmitted    bool             `json:"issubmitted"`
	QuestionReport []QuestionReport `json:"questionreport"`
}

type ExamReport struct {
	ExamID         string          `json:"examid"`
	AttemptReports []AttemptReport `json:"attemptreports"`
}
