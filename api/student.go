package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func AddStudent(getStudent models.Student) (string, error) {
	uuid, _ := uuid.New()
	uuidstring := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	student := models.Student{
		StudentID:     uuidstring,
		StudentName:   getStudent.StudentName,
		College:       getStudent.College,
		Username:      getStudent.Username,
		Email:         getStudent.Email,
		Password:      getStudent.Password,
		Courses:       getStudent.Courses,
		CourseReports: []models.CourseReport{},
	}

	file, _ := json.MarshalIndent(student, "", " ")

	path := fmt.Sprintf("%s%s%s", "database/Student/", uuidstring, ".json")
	err := ioutil.WriteFile(path, file, 0644)
	return path, err
}

func GetStudent(path string) models.Student {

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	student := models.Student{}

	if err := json.Unmarshal([]byte(file), &student); err != nil {
		fmt.Println(err)
	}

	return student
}
