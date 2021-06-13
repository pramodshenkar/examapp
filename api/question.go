package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
)

func GetQuestionsIDsByCourseID(courseid, examid string) ([]string, error) {

	path := fmt.Sprintf("%s%s%s%s%s", "database/Course/", courseid, "/", examid, ".json")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	exam := models.Exam{}

	if err := json.Unmarshal([]byte(file), &exam); err != nil {
		return []string{}, err
	}

	return exam.Questions, nil
}

func GetQuestionsByQuestionID(courseid, questionid string) (models.Question, error) {

	path := fmt.Sprintf("%s%s%s%s%s%s%s%s", "database/Course/", courseid, "/", "questionbank_", courseid, "/", questionid, ".json")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Cannot read question file", err)
		return models.Question{}, err
	}
	question := models.Question{}

	if err := json.Unmarshal([]byte(file), &question); err != nil {

		fmt.Println("Cannot Unmarshall question file", err)
		return models.Question{}, err
	}
	return question, nil
}
