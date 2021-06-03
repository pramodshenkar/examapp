package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Student struct {
	StudentID    string         `json:"studentid,omitempty"`
	StudentName  string         `json:"name"`
	College      string         `json:"college"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Course       []string       `json:"course"`
	CourseReport []CourseReport `json:"reports,omitempty"`
}

func AddStudent(student Student) (string, error) {
	uuid, _ := uuid.New()
	uuidstring := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	student.StudentID = uuidstring
	student.CourseReport = []CourseReport{}

	file, _ := json.MarshalIndent(student, "", " ")

	path := fmt.Sprintf("%s%s%s", "database/Student/", uuidstring, ".json")
	err := ioutil.WriteFile(path, file, 0644)
	return path, err
}

func GetStudent(path string) Student {

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	student := Student{}

	if err := json.Unmarshal([]byte(file), &student); err != nil {
		fmt.Println(err)
	}

	return student
}
