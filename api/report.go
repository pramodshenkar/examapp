package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pramodshenkar/examapp/models"
)

func GetReport(userid string, courseid string, examid string) (models.ExamReport, error) {

	filename := fmt.Sprintf("%s%s%s%s%s", userid, "_", courseid, "_", examid)

	fmt.Println(filename)

	path := fmt.Sprintf("%s%s%s", "database/Report/", filename, ".json")

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Problem while reading file", err)
		return models.ExamReport{}, err
	}

	examReport := models.ExamReport{}

	if err := json.Unmarshal([]byte(file), &examReport); err != nil {
		fmt.Println("Problem while unmarshalling file")
		return models.ExamReport{}, err
	}

	return examReport, nil
}
