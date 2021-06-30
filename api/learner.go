package api

import (
	"context"
	"fmt"

	"github.com/pramodshenkar/examapp/connectionHelper"
	"github.com/pramodshenkar/examapp/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddLearner(learner models.Learner) (*mongo.InsertOneResult, error) {
	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.LEARNER)

	result, err := collection.InsertOne(context.TODO(), learner)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func GetAllLearners() ([]models.Learner, error) {

	learners := []models.Learner{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return []models.Learner{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.LEARNER)

	filter := bson.D{{}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Learner{}, err
	}

	for cur.Next(context.TODO()) {
		learner := models.Learner{}

		err := cur.Decode(&learner)
		if err != nil {
			return []models.Learner{}, err
		}

		learners = append(learners, learner)
	}

	cur.Close(context.TODO())

	if len(learners) == 0 {
		return []models.Learner{}, mongo.ErrNoDocuments
	}

	return learners, nil
}

func GetLearner(username string) (models.Learner, error) {
	learner := models.Learner{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return models.Learner{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.LEARNER)

	// objectID, _ := primitive.ObjectIDFromHex(learnerid)
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	err = collection.FindOne(context.TODO(), filter).Decode(&learner)
	if err != nil {
		return models.Learner{}, err
	}

	return learner, nil
}

func DeleteLearner(learnerid string) (*mongo.DeleteResult, error) {
	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.LEARNER)

	objectID, _ := primitive.ObjectIDFromHex(learnerid)

	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func UpdateLearner(learner models.Learner) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "username", Value: learner.Username}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		{Key: "learnername", Value: learner.LearnerName},
		{Key: "username", Value: learner.Username},
		{Key: "password", Value: learner.Password},
		{Key: "college", Value: learner.College},
		{Key: "email", Value: learner.Email},
		{Key: "courses", Value: learner.Courses},
		{Key: "examreport", Value: learner.Report},
	}}}

	fmt.Println(filter, updater)

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.LEARNER)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}
