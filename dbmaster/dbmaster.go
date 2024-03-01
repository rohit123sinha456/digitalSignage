package dbmaster

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func ConnectDB() *mongo.Client {
	const uri = "mongodb://localhost:27017"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if client != nil {
		log.Printf("Client successfully Created")
	}
	log.Printf("Successfully Connected to database")
	return client
}

func CreateUser(client *mongo.Client, newUser DataModel.User) string {
	if client == nil {
		log.Printf("Client is Null")
	}
	coll := client.Database("user").Collection("userData")
	// userID := uuid.NewString()
	// newUser := DataModel.User{Name: username, UserID: userID}
	_, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	log.Printf("Created User")

	return newUser.UserID
}

func GetUser(client *mongo.Client, userId string) {
	var result DataModel.User
	coll := client.Database("user").Collection("userData")
	filter := bson.D{{"user_id", userId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
