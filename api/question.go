package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
)

func GetQuestionsByExamID(courseid, examid string) ([]models.Question, error) {
	// path := fmt.Sprintf("%s%s%s%s%s", "database/Course/", courseid, "/", examid, "/*")

	// files, err := filepath.Glob(path)
	// if err != nil {
	// return []models.Question{}, err
	// }

	// var questions []models.Question

	// for _, file := range files {
	// 	file, err := ioutil.ReadFile(file)

	// 	if err != nil {
	// 		continue
	// 	}

	// 	question := models.Question{}

	// 	if err := json.Unmarshal([]byte(file), &question); err != nil {
	// 		continue
	// 	}

	// 	questions = append(questions, question)
	// }
	// fmt.Println(questions)
	// return questions, nil
	return []models.Question{}, nil
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
