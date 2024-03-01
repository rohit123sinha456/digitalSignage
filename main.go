package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

func main() {
	client := dbmaster.ConnectDB()

	userID := uuid.NewString()
	newUser := DataModel.User{Name: "Sonia Di", UserID: userID}
	userid := dbmaster.CreateUser(client, newUser)
	fmt.Printf(userid)

	// dbmaster.GetUser(client, "55aa7c7b-96e4-44d4-a188-8520f104eac4")

	//Create Playlist
	playlistID := uuid.NewString()
	var ptype DataModel.PlaylistType = 0
	playlist := DataModel.Playlist{
		ID:       playlistID,
		DeviceId: "xx",
		PType:    ptype,
		DisplayBlock: []DataModel.DisplayBlock{
			{BlockName: "aa", Imagelist: []DataModel.ImageBlock{{Imagetype: "JPEG", Image: "hdhdhd", DisplayTime: 10}}},
		},
	}
	dbmaster.CreatePlaylist(client, userID, playlist)
}
