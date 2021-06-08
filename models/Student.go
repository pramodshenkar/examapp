package models

type Student struct {
	StudentID     string         `json:"studentid"`
	StudentName   string         `json:"name" binding:"required"`
	Username      string         `json:"username" binding:"required"`
	College       string         `json:"college"`
	Email         string         `json:"email" binding:"required,email"`
	Password      string         `json:"password" binding:"required"`
	Courses       []string       `json:"courses"`
	CourseReports []CourseReport `json:"reports"`
}
