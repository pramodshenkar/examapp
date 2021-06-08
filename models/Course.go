package models

type Course struct {
	CourseID   string   `json:"courseid"`
	CourseName string   `json:"coursename"`
	Exams      []string `json:"exams,omitempty"`
}
