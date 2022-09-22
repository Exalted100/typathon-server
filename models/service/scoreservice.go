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

type HighestScores struct{
	Sprint []*entity.Score `json:"sprint"`
	MiddleDistance []*entity.Score `json:"middleDistance"`
	Marathon []*entity.Score `json:"marathon"`
}

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
func (scoreservice ScoreService) FindUserHighestScores(parameter string, value string) (HighestScores, error) {
	opts := options.Find().SetSort(bson.M{"score": -1}).SetLimit(5)
	sprintCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "sprint", parameter: value}, opts)
	if err != nil {
		fmt.Println(sprintCursor)
		return HighestScores{}, err
	}
	defer sprintCursor.Close(context.TODO())
	var sprintScores []*(entity.Score)

	for sprintCursor.Next(context.TODO()) {
		var score *(entity.Score)
		sprintCursor.Decode(&score)
		sprintScores = append(sprintScores, score)
	}
	if err := sprintCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	middleDistanceCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "middleDistance", parameter: value}, opts)
	if err != nil {
		fmt.Println(middleDistanceCursor)
		return HighestScores{}, err
	}
	defer middleDistanceCursor.Close(context.TODO())
	var middleDistanceScores []*(entity.Score)

	for middleDistanceCursor.Next(context.TODO()) {
		var score *(entity.Score)
		middleDistanceCursor.Decode(&score)
		middleDistanceScores = append(middleDistanceScores, score)
	}
	if err := middleDistanceCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	marathonCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "marathon", parameter: value}, opts)
	if err != nil {
		fmt.Println(marathonCursor)
		return HighestScores{}, err
	}
	defer marathonCursor.Close(context.TODO())
	var marathonScores []*(entity.Score)

	for marathonCursor.Next(context.TODO()) {
		var score *(entity.Score)
		marathonCursor.Decode(&score)
		marathonScores = append(marathonScores, score)
	}
	if err := marathonCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	return HighestScores{
		Sprint: sprintScores,
		MiddleDistance: middleDistanceScores,
		Marathon: marathonScores,
	}, nil
}

//Find highest user scores
func (scoreservice ScoreService) FindHighestScores() (HighestScores, error) {
	opts := options.Find().SetSort(bson.M{"score": -1}).SetLimit(5)
	sprintCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "sprint"}, opts)
	if err != nil {
		fmt.Println(sprintCursor)
		return HighestScores{}, err
	}
	defer sprintCursor.Close(context.TODO())
	var sprintScores []*(entity.Score)

	for sprintCursor.Next(context.TODO()) {
		var score *(entity.Score)
		sprintCursor.Decode(&score)
		sprintScores = append(sprintScores, score)
	}
	if err := sprintCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	middleDistanceCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "middleDistance"}, opts)
	if err != nil {
		fmt.Println(middleDistanceCursor)
		return HighestScores{}, err
	}
	defer middleDistanceCursor.Close(context.TODO())
	var middleDistanceScores []*(entity.Score)

	for middleDistanceCursor.Next(context.TODO()) {
		var score *(entity.Score)
		middleDistanceCursor.Decode(&score)
		middleDistanceScores = append(middleDistanceScores, score)
	}
	if err := middleDistanceCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	marathonCursor, err := db.ScoresCollection.Find(context.TODO(), bson.M{"mode": "marathon"}, opts)
	if err != nil {
		fmt.Println(marathonCursor)
		return HighestScores{}, err
	}
	defer marathonCursor.Close(context.TODO())
	var marathonScores []*(entity.Score)

	for marathonCursor.Next(context.TODO()) {
		var score *(entity.Score)
		marathonCursor.Decode(&score)
		marathonScores = append(marathonScores, score)
	}
	if err := marathonCursor.Err(); err != nil {
		fmt.Println("There was an error while retrieving highest scores", err)
	}

	return HighestScores{
		Sprint: sprintScores,
		MiddleDistance: middleDistanceScores,
		Marathon: marathonScores,
	}, nil
}
