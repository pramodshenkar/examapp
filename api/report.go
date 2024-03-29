package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
)

func GetReport(userid string, courseid string, examid string, attemptCount int) (models.ExamReport, error) {

	filename := fmt.Sprintf("%s%s%s%s%s", userid, "_", courseid, "_", examid)

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

	// fmt.Println(questionReports)

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
		GivenAnswer: []string{},
		Marks:       0,
	}
	return questionReport
}

func UpdateReportForEndExam(userid string, courseid string, examid string) bool {

	filename := fmt.Sprintf("%s%s%s%s%s", userid, "_", courseid, "_", examid)

	path := fmt.Sprintf("%s%s%s", "database/Report/", filename, ".json")

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Problem while reading file", err)
		return false
	}

	examReport := models.ExamReport{}

	if err := json.Unmarshal([]byte(file), &examReport); err != nil {
		fmt.Println("Problem while unmarshalling file")
		return false
	}

	isUpdated := false
	var attemptReports []models.AttemptReport
	for _, attemptReport := range examReport.AttemptReports {

		if !isUpdated {
			if !attemptReport.IsSubmitted {
				attemptReport.IsSubmitted = true
				isUpdated = true
			}
		}
		attemptReports = append(attemptReports, attemptReport)
	}
	examReport.AttemptReports = attemptReports

	if isUpdated {

		// fmt.Println(examReport)

		file, err = json.MarshalIndent(examReport, "", " ")

		if err != nil {
			fmt.Println(err)
			return false
		}

		err = ioutil.WriteFile(path, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
		return err == nil
	}

	return isUpdated
}

func UpdateReportForSubmitAnswer(userid, courseid, examid, questionid string, answerid []string) (bool, error) {

	// fmt.Println("userid", userid, "\ncourseid", courseid, "\nexamid", examid, "\nquestionid", questionid, "\nanswerid", answerid)
	filename := fmt.Sprintf("%s%s%s%s%s", userid, "_", courseid, "_", examid)

	path := fmt.Sprintf("%s%s%s", "database/Report/", filename, ".json")

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Problem while reading file", err)
		return false, err
	}

	examReport := models.ExamReport{}

	if err := json.Unmarshal([]byte(file), &examReport); err != nil {
		fmt.Println("Problem while unmarshalling file")
		return false, err
	}

	isUpdated := false
	var attemptReports []models.AttemptReport
	for _, attemptReport := range examReport.AttemptReports {

		// fmt.Println("----------------------------------------------------------------")
		// fmt.Println(i, "attemptReport : ", attemptReport, "\nisUpdated : ", isUpdated)
		if !isUpdated {
			if !attemptReport.IsSubmitted {

				var questionReports []models.QuestionReport

				for _, questionReport := range attemptReport.QuestionReport {

					// fmt.Println("	- ", j, "question : ", questionReport)

					if questionReport.QuestionID == questionid {

						// fmt.Println("		- ", "if ", questionReport.QuestionID, "==", questionid)
						// fmt.Println("			- ", "lets change data")

						questionReport.IsAnswered = true
						questionReport.GivenAnswer = answerid

						questionReport.Marks = GetMarks(courseid, questionReport.QuestionID, answerid)
						// fmt.Println("				- ", "assigning marks", questionReport.Marks)

						isUpdated = true
						// }
					}
					questionReports = append(questionReports, questionReport)
				}
				attemptReport.QuestionReport = questionReports
			}
		}
		attemptReports = append(attemptReports, attemptReport)
	}
	examReport.AttemptReports = attemptReports

	if isUpdated {

		file, err = json.MarshalIndent(examReport, "", " ")
		if err != nil {
			fmt.Println(err)
			return false, err
		}

		err = ioutil.WriteFile(path, file, 0644)
		if err != nil {
			fmt.Println(err)
			return false, err

		}
		return err == nil, err
	}

	return isUpdated, nil
}

func GetMarks(courseid, questionid string, givenAnswers []string) int {
	question, err := GetQuestionsByQuestionID(courseid, questionid)

	if err != nil {
		fmt.Println("can't get question", err)
		return 0
	}

	isCorrect := false

	for _, correctAnswer := range question.Answer {
		isCorrect = false
		for _, givenAnswer := range givenAnswers {
			if givenAnswer == correctAnswer.OptionId {
				isCorrect = true
				break
			}
		}
		if !isCorrect {
			break
		}
	}

	if isCorrect {
		return question.Marks
	}

	return 0
}
