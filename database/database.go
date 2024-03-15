package database

import (
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBSet() *mongo.Client {
	Client := dbmaster.ConnectDB()
	return Client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("user").Collection("userData")
	return collection
}
