package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

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

func GetStudentByUsername(username string) (models.Student, error) {

	StudentID, err := GetStudentFile(username)

	if err != nil {
		return models.Student{}, err
	}

	if StudentID == "" {
		fmt.Println("NO FILE FOUND")
		return models.Student{}, nil
	}

	path := fmt.Sprintf("%s%s%s", "database/Student/", StudentID, ".json")

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return models.Student{}, err
	}

	student := models.Student{}

	if err := json.Unmarshal([]byte(file), &student); err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func GetStudentFile(username string) (string, error) {
	files, err := filepath.Glob("./database/Student/*")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		file, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Println(err)
		}

		student := models.Student{}

		if err := json.Unmarshal([]byte(file), &student); err != nil {
			continue
		}

		if username == student.Username {
			return student.StudentID, err

		}
	}
	return "", err
}
