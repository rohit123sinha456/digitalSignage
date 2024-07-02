package dbmaster

import (
	"context"
	"log"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateContent(ctx context.Context, client *mongo.Client, userID string, content DataModel.Content) (string, error) {
	content.ID = primitive.NewObjectID()
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("content")
	result, inserterr := coll.InsertOne(ctx, content)
	if inserterr != nil {
		return "", inserterr
	}
	log.Printf("Content Created")
	idstring := result.InsertedID.(primitive.ObjectID).Hex()
	return idstring, nil
}

func ReadContent(ctx context.Context, client *mongo.Client, userID string) ([]DataModel.Content, error) {
	var contentarray []DataModel.Content
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("content")
	filter := bson.D{{}}
	curr, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		var result DataModel.Content
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}
		contentarray = append(contentarray, result)
	}
	defer curr.Close(ctx)
	return contentarray, nil
}

func ReadOneContent(ctx context.Context, client *mongo.Client, userID string, contentId string) (DataModel.Content, error) {
	var result DataModel.Content
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("content")
	objectId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		return result, err
	}
	filter := bson.D{{"_id", objectId}}
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteContent(ctx context.Context, client *mongo.Client, userID string, contentId string) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	contentCollection := client.Database(userSystemname).Collection("content")
	userdBname := common.ExtractUserSystemIdentifier(userID)
	plalistCollection := client.Database(userdBname).Collection("playlist")

	objectId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	result, err := contentCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("Number of documents deleted from Content: %d\n", result.DeletedCount)
	//db.playlist.updateMany({"deviceblock.displayblock.imagelist":{$elemMatch:{imageid: ObjectId('6683a8bca0cf1a28f6edddd3')}}},{$pull:{"deviceblock.$[].displayblock.$[].imagelist":{imageid:ObjectId('6683a8bca0cf1a28f6edddd3')}}})
	contentfilter := bson.M{
        "deviceblock.displayblock.imagelist": bson.M{
            "$elemMatch": bson.M{
                "imageid": objectId,
            },
        },
    }
    contentupdate := bson.M{
        "$pull": bson.M{
            "deviceblock.$[].displayblock.$[].imagelist": bson.M{
                "imageid": objectId,
            },
        },
    }

    // Perform the update operation
    playlistdeleteresult, err := plalistCollection.UpdateMany(ctx, contentfilter, contentupdate)
    if err != nil {
        return err
    }
	log.Printf("Number of Images deleted from Plalist: %d\n", playlistdeleteresult.ModifiedCount)
	return nil
}
// func UpdateContent() {}
