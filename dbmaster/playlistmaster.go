package dbmaster

import (
	"context"
	"log"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePlaylist(client *mongo.Client, userID string, playlist DataModel.Playlist) {
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("playlist")
	_, err := coll.InsertOne(context.TODO(), playlist)
	if err != nil {
		panic(err)
	}
	log.Printf("Created User")
}
