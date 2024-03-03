package dbmaster

import (
	"context"
	"log"

	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDevice(ctx context.Context, client *mongo.Client, userid string, devices []DataModel.Device) (string, error) {
	var uid string
	uid = userid
	log.Printf("jhfhifkdf")
	log.Printf("%s", userid)
	var result DataModel.User
	coll := client.Database("user").Collection("userData")
	filter := bson.D{{"userid", userid}}
	for _, device := range devices {
		result.Devices = append(result.Devices, device)
	}
	update := bson.D{{"$set", bson.D{{"devices", result.Devices}}}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return "nil", err
	}

	return uid, nil

}
