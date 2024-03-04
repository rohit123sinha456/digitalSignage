package dbmaster

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/rabbitqueue"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePlaylist(ctx context.Context, client *mongo.Client, userID string, playlist DataModel.Playlist) (string, error) {
	_, err := GetUser(client, userID)
	if err != nil {
		return "nil", err
	}
	playlistid := uuid.NewString()
	playlist.ID = playlistid
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("playlist")
	_, inserterr := coll.InsertOne(ctx, playlist)
	if inserterr != nil {
		return "nil", inserterr
	}
	log.Printf("Created User Playlist")
	return playlistid, nil
}

func GetPlaylist(ctx context.Context, client *mongo.Client, userID string, playlistID string) (DataModel.Playlist, error) {
	var result DataModel.Playlist
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	filter := bson.D{{"id", playlistID}}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func PlayPlaylist(ctx context.Context, client *mongo.Client, userID string, playlistid string) error {
	_, err := GetUser(client, userID)
	if err != nil {
		return err
	}
	uservHostname := common.CreatevHostName(userID)
	userdsystemname := common.ExtractUserSystemIdentifier(userID)
	playlist, err := GetPlaylist(ctx, client, userID, playlistid)
	rabbitqueue.Connect(userdsystemname, "password", uservHostname)
	err = rabbitqueue.PublishMessage(ctx, playlist, uservHostname)
	if err != nil {
		return err
	}
	return nil

}
