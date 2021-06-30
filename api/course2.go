package api

import (
	"context"
	"fmt"

	"github.com/pramodshenkar/examapp/connectionHelper"
	"github.com/pramodshenkar/examapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllCourses() ([]models.Course, error) {

	courses := []models.Course{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return []models.Course{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.COURSE)

	filter := bson.D{{}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Course{}, err
	}

	for cur.Next(context.TODO()) {
		course := models.Course{}

		// fmt.Println(cur.Current)
		// fmt.Println(course)

		err := cur.Decode(&course)
		if err != nil {
			fmt.Println("cant find Record")

			return []models.Course{}, err
		}

		courses = append(courses, course)
	}

	cur.Close(context.TODO())

	if len(courses) == 0 {
		return []models.Course{}, mongo.ErrNoDocuments
	}

	return courses, nil
}

func AddCourse(course models.Course) (*mongo.InsertOneResult, error) {
	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.COURSE)

	result, err := collection.InsertOne(context.TODO(), course)
	if err != nil {
		return nil, err
	}

	return result, nil

}
