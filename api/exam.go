package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pramodshenkar/examapp/models"
)

func GetExams(courseid string) ([]models.Exam, error) {
	path := fmt.Sprintf("%s%s%s", "database/Course/", courseid, "/*")

	files, err := filepath.Glob(path)
	if err != nil {
		return []models.Exam{}, err
	}

	var exams []models.Exam

	for _, file := range files {
		file, err := ioutil.ReadFile(file)

		if err != nil {
			continue
		}

		data := models.Exam{}

		if err := json.Unmarshal([]byte(file), &data); err != nil {
			continue
		}

		exam := models.Exam{
			ExamID:    data.ExamID,
			ExamName:  data.ExamName,
			Questions: nil,
		}

		exams = append(exams, exam)
	}
	fmt.Println(exams)
	return exams, nil
}

func GetExamsByID(courseid, examid string) (models.Exam, error) {

	path := fmt.Sprintf("%s%s%s%s%s", "database/Course/", courseid, "/", examid, ".json")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Cannot read exam file", err)
		return models.Exam{}, err
	}
	exam := models.Exam{}

	if err := json.Unmarshal([]byte(file), &exam); err != nil {

		fmt.Println("Cannot Unmarshall exam file", err)
		return models.Exam{}, err
	}
	return exam, nil
}
