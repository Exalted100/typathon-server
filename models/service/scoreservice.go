package service

import (
	"context"
	"fmt"
	"typathon/models/db"
	"typathon/models/entity"

	"go.mongodb.org/mongo-driver/mongo/options"
	"labix.org/v2/mgo/bson"
)

//Userservice is to handle user relation db query
type ScoreService struct{}

//Create is to register new score
func (scoreservice ScoreService) Create(score *(entity.Score)) error {
	_, err := db.ScoresCollection.InsertOne(context.TODO(), score)

	if err != nil {
		return err
	}
	return nil
}

//Find score
func (scoreservice ScoreService) FindOne(parameter string, value string) (*entity.Score, error) {
	var score (*entity.Score)
	err := db.ScoresCollection.FindOne(context.TODO(), bson.M{parameter: value}).Decode(&score)

	if err != nil {
		return nil, err
	}
	return score, nil
}

//Find all user scores
func (scoreservice ScoreService) FindAll(parameter string, value string) ([]*entity.Score, error) {
	cursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{parameter: value})
	var scores []*(entity.Score)

	for cursor.Next(context.TODO()) {
		var score *(entity.Score)
		cursor.Decode(&score)
		scores = append(scores, score)
	}

	if err != nil {
		return nil, err
	}
	return scores, nil
}

//Find highest user scores
func (scoreservice ScoreService) FindHighestScores() ([]*entity.Score, error) {
	opts := options.Find().SetSort(bson.M{"score": -1}).SetLimit(10)
	cursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		fmt.Println(cursor)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var scores []*(entity.Score)
	fmt.Println("reached here")

	for cursor.Next(context.TODO()) {
		fmt.Println("loop running")
		var score *(entity.Score)
		cursor.Decode(&score)
		scores = append(scores, score)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	return scores, nil
}
