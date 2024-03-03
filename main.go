package main

import (
	"github.com/rohit123sinha456/digitalSignage/router"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// client := dbmaster.ConnectDB()
	// objectStroreClient := objectstore.ConnectObjectStore()

	// // Create Users
	// userID := uuid.NewString()
	// newUser := DataModel.User{Name: "Rohit Sinha", UserID: userID}
	// err := dbmaster.CreateUser(client, objectStroreClient,newUser)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf(userID)
	// log.Println("Successfully User Created")

	// Publish Message
	// userID := "008be3f4-7dcb-4b75-b322-403f6cb1d9ab"
	// newUser := DataModel.User{Name: "Rohit Sinha", UserID: userID}

	// uservhostname := common.CreatevHostName(userID)
	// userdsystemname := common.ExtractUserSystemIdentifier(userID)
	// rabbitqueue.Connect(userdsystemname, "password", uservhostname)
	// rabbitqueue.PublishMessage(ctx, newUser, uservhostname)

	// Add Bucket to Object STore
	// userbucketname := common.CreateBucketName(userID)
	// objectstore.CreateBucket(ctx, objectStroreClient, userbucketname)

	// dbmaster.GetUser(client, "55aa7c7b-96e4-44d4-a188-8520f104eac4")

	//Create Playlist
	// playlistID := uuid.NewString()
	// var ptype DataModel.PlaylistType = 0
	// playlist := DataModel.Playlist{
	// 	ID:       playlistID,
	// 	DeviceId: "xx",
	// 	PType:    ptype,
	// 	DisplayBlock: []DataModel.DisplayBlock{
	// 		{BlockName: "aa", Imagelist: []DataModel.ImageBlock{{Imagetype: "JPEG", Image: "hdhdhd", DisplayTime: 10}}},
	// 	},
	// }
	// dbmaster.CreatePlaylist(client, userID, playlist)
	router.SetupRouter()
	router.UserRouter()
	router.DeviceRouter()
	// Listen and Server in 0.0.0.0:8080
	router.R.Run(":8080")
}
