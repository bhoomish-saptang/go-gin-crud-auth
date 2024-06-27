package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-gin/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection
var CollectionAuth *mongo.Collection

func ConnectMongodb() {
	db_name := os.Getenv("DB_NAME")
	collection_name_auth := os.Getenv("COLLECTION_NAME_AUTH")
	collection_name := os.Getenv("COLLECTION_NAME")
	var err error

	clientOptions := options.Client().ApplyURI(os.Getenv("URI"))
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Unable to MongoDb client: ", err.Error())
	}

	collection = client.Database(db_name).Collection(collection_name)
	CollectionAuth = client.Database(db_name).Collection(collection_name_auth)

}

func InsertUserData(data constants.User) {
	insertResult, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err.Error())
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)
}

func DeleteUserData(data constants.User) {
	_, err := collection.DeleteOne(context.Background(), data)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err.Error())
	}

}

func FindUserDetailsByID(id string) (constants.User, error) {
	var findedUser constants.User
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&findedUser)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err)
		return findedUser, err
	}
	return findedUser, nil
}

func DeleteUserDetailsByID(id string) (constants.User, error) {
	var findedUser constants.User
	filter := bson.M{"id": id}
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&findedUser)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err)
		return findedUser, err
	}
	return findedUser, nil

}

func UpdateUserDetailsByID(id string, userDetails constants.User) (constants.User, error) {

	filter := bson.M{"id": id}
	//var data constants.User
	result, err := collection.ReplaceOne(context.Background(), filter, userDetails)
	if err != nil {
		return constants.User{}, err
	}

	fmt.Println(".Printing a upserted ID", result.UpsertedID)
	return userDetails, nil
}

func GetAllUserDetails() ([]constants.User, error) {

	AllUserDetails := []constants.User{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return AllUserDetails, err
	}
	for result.Next(context.Background()) {
		var User constants.User
		result.Decode(&User)
		AllUserDetails = append(AllUserDetails, User)
	}
	return AllUserDetails, nil

}
