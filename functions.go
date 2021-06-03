package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func CreateStudent() (string, string, error) {
	uuid, _ := uuid.New()
	uuidstring := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	data := Student{
		StudentID:    uuidstring,
		StudentName:  "Pramod",
		College:      "AVCOE",
		Email:        "pramod@mail.com",
		Password:     "abc123",
		Course:       nil,
		CourseReport: nil,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	path := fmt.Sprintf("%s%s%s", "database/Student/", uuidstring, ".json")
	err := ioutil.WriteFile(path, file, 0644)
	return data.StudentName, path, err
}
