package dbmaster

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/rabbitqueue"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreatePlaylistPaylaodRequest struct {
	ID       primitive.ObjectID `bson:"_id"`
	ScreenID primitive.ObjectID `bson:"screenid"`
}

func checkifscreenexists(ctx context.Context, coll *mongo.Collection, contentlist DataModel.DeviceBlock) {
	log.Printf("Screen Check")
}
func checkifimageexists(ctx context.Context, coll *mongo.Collection, contentlist DataModel.ImageBlock) {
	log.Printf("Image Check")
}

func CreatePlaylist(ctx context.Context, client *mongo.Client, userID string, playlist DataModel.Playlist) (DataModel.Playlist, error) {
	playlist.ID = primitive.NewObjectID()
	_, err := GetUser(client, userID)
	if err != nil {
		return playlist, err
	}
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("playlist")
	_, inserterr := coll.InsertOne(ctx, playlist)
	if inserterr != nil {
		return playlist, inserterr
	}
	log.Printf("Created User Playlist")
	return playlist, nil
}

func UpdatePlaylist(ctx context.Context, client *mongo.Client, userID string, playlistID string, updatejson DataModel.UpdatePlaylistRequest) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	objectId, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return err
	}
	updateDoc := bson.M{}
	if updatejson.Name != "" {
		updateDoc["playlistname"] = updatejson.Name
	}
	if len(updatejson.DeviceBlock) != 0 {
		updateDoc["deviceblock"] = updatejson.DeviceBlock
		fmt.Printf("%v\n", updatejson.DeviceBlock)

	}
	filter := bson.D{{"_id", objectId}}
	update := bson.M{"$set": updateDoc}
	result, updateerr := coll.UpdateOne(ctx, filter, update)
	if updateerr != nil {
		return updateerr
	}
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	return nil
}
func GetPlaylist(ctx context.Context, client *mongo.Client, userID string, playlistID string) (DataModel.Playlist, error) {
	var result DataModel.Playlist
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	objectId, err := primitive.ObjectIDFromHex(playlistID)
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

func ReadPlaylist(ctx context.Context, client *mongo.Client, userID string) ([]DataModel.Playlist, error) {
	var contentlistarray []DataModel.Playlist
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	filter := bson.D{{}}
	curr, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		var result DataModel.Playlist
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}
		contentlistarray = append(contentlistarray, result)
	}
	defer curr.Close(ctx)
	return contentlistarray, nil
}

func PlayPlaylist(ctx context.Context, client *mongo.Client, userID string, playlistid string) error {
	_, err := GetUser(client, userID)
	if err != nil {
		return err
	}
	uservHostname := common.CreatevHostName(userID)
	userdsystemname := common.ExtractUserSystemIdentifier(userID)
	playlist, err := GetPlaylist(ctx, client, userID, playlistid)
	if reflect.ValueOf(playlist).IsZero() == true {
		return errors.New(" Requested Playlist Doesnt Exists. Please check User ID and Playlist ID Mapping")
	}
	rabbitqueue.Connect(userdsystemname, "password", uservHostname)
	err = rabbitqueue.PublishMessage(ctx, playlist, uservHostname)
	if err != nil {
		return err
	}
	return nil

}
