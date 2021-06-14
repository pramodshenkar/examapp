package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func GenerateStudentID() string {
	uuid, _ := uuid.New()
	StudentID := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return StudentID
}

func AddStudent(studentid string, getStudent models.Student) bool {

	student := models.Student{
		StudentID:   studentid,
		StudentName: getStudent.StudentName,
		College:     getStudent.College,
		Username:    getStudent.Username,
		Email:       getStudent.Email,
		Password:    getStudent.Password,
		Courses:     getStudent.Courses,
	}

	file, err := json.MarshalIndent(student, "", " ")

	if err != nil {
		return false
	}

	path := fmt.Sprintf("%s%s%s", "database/Student/", studentid, ".json")
	err = ioutil.WriteFile(path, file, 0644)

	return err == nil
}

func GetStudent(username string) (models.Student, error) {

	studentCredentials, err := GetStudentCredentials(username)

	if err != nil {
		return models.Student{}, err
	}

	if studentCredentials.Username == "" {
		return models.Student{}, err
	}

	path := fmt.Sprintf("%s%s%s", "database/Student/", studentCredentials.StudentID, ".json")

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

func AddCredentials(studentid string, student models.Student) bool {

	allcredentials, err := GetAllCredentials()

	if err != nil {
		return false
	}

	credentials := models.Credentials{
		StudentID: studentid,
		Username:  student.Username,
		Password:  student.Password,
	}

	allcredentials = append(allcredentials, credentials)

	fmt.Println(allcredentials)

	credentialFile := struct {
		Credentials []models.Credentials `json:"credentials"`
	}{
		Credentials: allcredentials,
	}

	file, _ := json.MarshalIndent(credentialFile, "", " ")

	err = ioutil.WriteFile("database/credentials.json", file, 0644)

	return err == nil
}

func GetAllCredentials() ([]models.Credentials, error) {

	file, err := ioutil.ReadFile("database/credentials.json")

	if err != nil {
		fmt.Println(err)
		return []models.Credentials{}, err
	}

	var credentials struct {
		Credentials []models.Credentials `json:"credentials"`
	}

	if err := json.Unmarshal([]byte(file), &credentials); err != nil {
		fmt.Println(err)
		return []models.Credentials{}, err
	}

	return credentials.Credentials, nil
}

func GetStudentCredentials(username string) (models.Credentials, error) {

	credentials, err := GetAllCredentials()
	if err != nil {
		return models.Credentials{}, err
	}

	for _, credential := range credentials {

		if credential.Username == username {
			return credential, nil
		}
	}
	return models.Credentials{}, err
}
