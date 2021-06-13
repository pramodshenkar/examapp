package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
)

func GetReport(userid string, courseid string, examid string, attemptCount int) (models.ExamReport, error) {

	filename := fmt.Sprintf("%s%s%s%s%s", userid, "_", courseid, "_", examid)

	fmt.Println(filename)

	path := fmt.Sprintf("%s%s%s", "database/Report/", filename, ".json")

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Problem while reading file", err)
		examReport, err := GenerateReport(path, courseid, examid, attemptCount)

		if err != nil {
			fmt.Println(err)
			return models.ExamReport{}, err
		}

		return examReport, nil
	}

	examReport := models.ExamReport{}

	if err := json.Unmarshal([]byte(file), &examReport); err != nil {
		fmt.Println("Problem while unmarshalling file")
		return models.ExamReport{}, err
	}

	return examReport, nil
}

func GenerateReport(path, courseid, examid string, attemptCount int) (models.ExamReport, error) {

	var attemptReports []models.AttemptReport
	for attemptNo := 1; attemptNo < attemptCount+1; attemptNo++ {
		attemptReport := GenerateAttemptReport(courseid, examid, attemptNo)
		attemptReports = append(attemptReports, attemptReport)
	}

	examReport := models.ExamReport{
		ExamID:         examid,
		AttemptReports: attemptReports,
	}

	file, err := json.MarshalIndent(examReport, "", " ")

	if err != nil {
		fmt.Println(err)
		return models.ExamReport{}, err
	}

	err = ioutil.WriteFile(path, file, 0644)

	if err != nil {
		return models.ExamReport{}, err
	}

	return examReport, nil

}

func GenerateAttemptReport(courseid, examid string, attemptNo int) models.AttemptReport {
	var questionReports []models.QuestionReport

	questions, err := GetQuestionsIDsByExamID(courseid, examid)

	if err != nil {
		fmt.Println("Error while getting Questions to generate report", err)
		return models.AttemptReport{}
	}
	for _, questionid := range questions {
		questionReport := GenerateQuestionReport(questionid)
		questionReports = append(questionReports, questionReport)
	}

	fmt.Println(questionReports)

	attemptReport := models.AttemptReport{
		AttemptNo:      attemptNo,
		IsSubmitted:    false,
		QuestionReport: questionReports,
	}

	return attemptReport
}

func GenerateQuestionReport(questionid string) models.QuestionReport {
	questionReport := models.QuestionReport{
		QuestionID:  questionid,
		IsAnswered:  false,
		GivenAnswer: "",
		Marks:       0,
	}
	return questionReport
}
