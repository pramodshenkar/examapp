package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pramodshenkar/examapp/models"
)

func GetCourses() ([]models.Course, error) {
	files, err := filepath.Glob("./database/Course/*")
	if err != nil {
		return []models.Course{}, err
	}

	var courses []models.Course

	for _, file := range files {
		file, err := ioutil.ReadFile(file)

		if err != nil {
			continue
		}

		course := models.Course{}

		if err := json.Unmarshal([]byte(file), &course); err != nil {
			continue
		}
		courses = append(courses, course)
	}
	fmt.Println(courses)
	return courses, nil
}

func GetCoursesByUsername(student models.Student) ([]models.Course, error) {
	// files, err := filepath.Glob("./database/Course/*")
	// if err != nil {
	// 	return []models.Course{}, err
	// }

	courseids := student.Courses

	fmt.Println(courseids)

	var courses []models.Course

	for _, courseid := range courseids {
		path := fmt.Sprintf("%s%s%s", "database/Course/", courseid, ".json")
		file, err := ioutil.ReadFile(path)

		if err != nil {
			continue
		}

		course := models.Course{}

		if err := json.Unmarshal([]byte(file), &course); err != nil {
			continue
		}
		courses = append(courses, course)
	}
	return courses, nil
}
