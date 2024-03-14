package dbmaster

import (
	"context"
	"errors"
	"log"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkifcontentexists(ctx context.Context, coll *mongo.Collection, contentlist DataModel.ContentList) error {
	log.Printf("Checking Contents")
	var objectIDs []primitive.ObjectID

	for _, contentblock := range contentlist.ContentList {
		log.Printf("%s", contentblock.Content)
		contentId := contentblock.Content
		objectId, err := primitive.ObjectIDFromHex(contentId)
		if err != nil {
			return err
		}
		objectIDs = append(objectIDs, objectId)
	}
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}
	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("Count is %d", count)
	if len(objectIDs) != int(count) {
		return errors.New("Count Mismatch Error")
	}
	return nil
}

func CreateContentList(ctx context.Context, client *mongo.Client, userID string, contentlist DataModel.ContentList) (string, error) {

	contentlist.ID = primitive.NewObjectID()
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("contentlist")
	contentCollection := client.Database(userdBname).Collection("content")
	err := checkifcontentexists(ctx, contentCollection, contentlist)
	if err != nil {
		return "", err
	}
	result, inserterr := coll.InsertOne(ctx, contentlist)
	if inserterr != nil {
		return "", inserterr
	}
	log.Printf("Contentlist Created")
	idstring := result.InsertedID.(primitive.ObjectID).Hex()
	return idstring, nil
}

func ReadContentList(ctx context.Context, client *mongo.Client, userID string) ([]DataModel.ContentList, error) {
	var contentlistarray []DataModel.ContentList
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("contentlist")
	filter := bson.D{{}}
	curr, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		var result DataModel.ContentList
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}
		contentlistarray = append(contentlistarray, result)
	}
	defer curr.Close(ctx)
	return contentlistarray, nil
}

// func UpdateContentList(){}
// func RemoveContentList(){}
