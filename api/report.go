package api

import "github.com/pramodshenkar/examapp/models"

func InitQuestionReport(questionid string) (models.QuestionReport, error) {

	questionReport := models.QuestionReport{
		QuestionID:  questionid,
		IsAnswered:  false,
		GivenAnswer: "",
		Marks:       0,
	}
	return questionReport, nil
}

func InitExamReport(examid string) (models.ExamReport, error) {

	var questionReports []models.QuestionReport

	examReport := models.ExamReport{
		ExamID:          examid,
		IsSubmitted:     false,
		QuestionReports: questionReports,
	}

	return examReport, nil
}

func InitCourseReport(courseid string) (models.CourseReport, error) {

	var examReports []models.ExamReport

	courseReport := models.CourseReport{
		CourseID:    courseid,
		ExamReports: examReports,
	}

	return courseReport, nil
}

func GenerateCourseReport(courses []string) ([]models.CourseReport, error) {

	var courseReports []models.CourseReport

	for _, course := range courses {
		courseReport, err := InitCourseReport(course)

		if err != nil {
			continue
		}
		courseReports = append(courseReports, courseReport)
	}

	return courseReports, nil
}

func GenerateExamReport(exams []string) ([]models.ExamReport, error) {

	var examReports []models.ExamReport

	for _, exam := range exams {
		examReport, err := InitExamReport(exam)

		if err != nil {
			continue
		}
		examReports = append(examReports, examReport)
	}

	return examReports, nil
}

func GenerateQuestionReport(questions []string) ([]models.QuestionReport, error) {

	var questionReports []models.QuestionReport

	for _, question := range questions {
		questionReport, err := InitQuestionReport(question)

		if err != nil {
			continue
		}
		questionReports = append(questionReports, questionReport)
	}

	return questionReports, nil
}
