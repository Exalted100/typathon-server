package db

import (
	"context"
	"fmt"
	"log"
	"typathon/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var UsersCollection *mongo.Collection
var ScoresCollection *mongo.Collection

func ConnectMongoDB(ConfigVariables config.ConfigType) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConfigVariables.MongodbUrl))
	if err != nil {
		fmt.Println("Could not connect to mongodb database", ConfigVariables.MongodbUrl)
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Println("Could not ping mongodb database", ConfigVariables.MongodbUrl)
		log.Fatal(err)
	}

	UsersCollection = client.Database("typathon").Collection("users")
	ScoresCollection = client.Database("typathon").Collection("scores")
}
