package models

type Course struct {
	CourseID   string   `json:"courseid"`
	CourseName string   `json:"coursename"`
	Exams      []string `json:"exams"`
}

type Exam struct {
	ExamID    string   `json:"examid"`
	ExamName  string   `json:"examname"`
	Attempts  int      `json:"attempts"`
	Questions []string `json:"questions,omitempty"`
}

type Question struct {
	QuestionID   string   `json:"questionid"`
	QuestionText string   `json:"questiontext"`
	Options      []Option `json:"options"`
	Answer       Option   `json:"answer"`
	Marks        int      `json:"marks"`
}

type Option struct {
	OptionId    string `json:"optionid"`
	OptionsText string `json:"optiontext"`
}
