package dbmaster

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/rabbitqueue"
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

func CreateUser(client *mongo.Client, newUser DataModel.User) error {
	userID := newUser.UserID
	coll := client.Database("user").Collection("userData")
	_, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}

	uservhostname := common.CreatevHostName(userID)
	userdsystemname := common.ExtractUserSystemIdentifier(userID)
	err = rabbitqueue.SetupUserandvHost(userdsystemname, uservhostname)
	log.Printf("Created User")
	if err != nil {
		return err
	}
	return nil
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

func AddUserDevice(client *mongo.Client, userId string) {
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
