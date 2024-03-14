package dbmaster

import (
	"context"
	"errors"
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

// This function gets a ID of ContentList and returns a list of content
func getcontentlistfromID(ctx context.Context, client *mongo.Client, userID string, contentListId string) ([]DataModel.ImageBlock, error) {
	var imageblocks []DataModel.ImageBlock
	contentList, err := ReadOneContentList(ctx, client, userID, contentListId)
	if err != nil {
		return imageblocks, err
	}
	for _, contentBlock := range contentList.ContentList {
		var x DataModel.ImageBlock
		x.Imagetype = contentBlock.Type
		tempcontent, err := ReadOneContent(ctx, client, userID, contentBlock.Content)
		if err != nil {
			return imageblocks, err
		}
		x.Image = tempcontent.Link
		x.DisplayTime = contentBlock.DisplayTime
		imageblocks = append(imageblocks, x)
	}

	return imageblocks, nil
}

// This function takes only the screen ID and construct the Playlist Payload by fetching from various databases
func createPlaylistPayload(ctx context.Context, client *mongo.Client, userID string, screen DataModel.Screen) (DataModel.Playlist, error) {
	var newPlaylist DataModel.Playlist
	var displayblockarr []DataModel.DisplayBlock
	newPlaylist.ID = primitive.NewObjectID()
	newPlaylist.DeviceId = screen.ID
	for _, screenblock := range screen.Screenblock {
		var displayblock DataModel.DisplayBlock
		displayblock.BlockName = screenblock.BlockName
		contentlistidstr := screenblock.ContentListID.Hex()
		imagelistfordisplayblock, err := getcontentlistfromID(ctx, client, userID, contentlistidstr)
		if err != nil {
			return newPlaylist, err
		}
		displayblock.Imagelist = imagelistfordisplayblock
		displayblockarr = append(displayblockarr, displayblock)
	}
	newPlaylist.DisplayBlock = displayblockarr
	return newPlaylist, nil

}

func CreatePlaylist(ctx context.Context, client *mongo.Client, userID string, screenid string) (DataModel.Playlist, error) {
	var playlist DataModel.Playlist
	_, err := GetUser(client, userID)
	if err != nil {
		return playlist, err
	}
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("playlist")
	screen, err := ReadOneScreen(ctx, client, userID, screenid)
	if err != nil {
		return playlist, err
	}
	playlist, err = createPlaylistPayload(ctx, client, userID, screen)
	if err != nil {
		return playlist, err
	}
	_, inserterr := coll.InsertOne(ctx, playlist)
	if inserterr != nil {
		return playlist, inserterr
	}
	log.Printf("Created User Playlist")
	return playlist, nil
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
