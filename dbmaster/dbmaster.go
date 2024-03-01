package dbmaster

import (
	"context"
	"log"

	"github.com/google/uuid"
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

func CreateUser(client *mongo.Client, username string) {
	if client == nil {
		log.Printf("Client is Null")
	}
	coll := client.Database("user").Collection("userData")
	userID = uuid.NewString()
	newUser := userModel.User{Name: username, UserID: userID}
	_, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	log.Printf("Created User")
	userdatabase := client.database(userID)
	if userdatabase.Name() != nil {
		log.Printf(userdatabase.Name())
	}
}
