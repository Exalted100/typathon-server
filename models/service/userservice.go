package service

import (
	"context"
	"errors"
	"typathon/models/db"
	"typathon/models/entity"

	"labix.org/v2/mgo/bson"
)

//Userservice is to handle user relation db query
type Userservice struct{}

//Create is to register new user
func (userservice Userservice) Create(user *(entity.User)) error {
	err := db.UsersCollection.FindOne(context.TODO(), bson.M{"email": user.Email})
	if err == nil {
		return errors.New("Account with this email Already Exists")
	}
	userNameErr := db.UsersCollection.FindOne(context.TODO(), bson.M{"userName": user.UserName})
	if userNameErr == nil {
		return errors.New("Account with this username Already Exists")
	}

	_, findErr := db.UsersCollection.InsertOne(context.TODO(), user)

	return findErr
}

// Delete a user from DB
func (userservice Userservice) Delete(email string) error {
	_, err := db.UsersCollection.DeleteOne(context.TODO(), bson.M{"email": email})
	return err
}

//Find user
func (userservice Userservice) Find(parameter string, value string) (*entity.User, error) {
	var user *(entity.User)
	err := db.UsersCollection.FindOne(context.TODO(), bson.M{parameter: value}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

//Find user from email
func (userservice Userservice) FindByEmail(email string) (*entity.User, error) {
	return userservice.Find("email", email)
}

func (userservice Userservice) UpdateOne(parameter string, value string, updateDoc bson.M) (error) {
	_, err := db.UsersCollection.UpdateOne(context.TODO(), bson.M{parameter: value}, updateDoc)
	return err
}