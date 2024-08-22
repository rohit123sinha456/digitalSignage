package dbmaster

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"
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
	now := time.Now()
	playlist.CreatedAt = &now
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
	now := time.Now()
	updateDoc["updatedAt"] = &now

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

func GetPlaylistwithSingleScreenData(ctx context.Context, client *mongo.Client, userID string, playlistID string,screenID string) ([]DataModel.Playlist, error) {
	var result []DataModel.Playlist
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	playlistobjectId, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return result, err
	}
	screenobjectId, err := primitive.ObjectIDFromHex(screenID)
	if err != nil {
		return result, err
	}

	// Build the aggregation pipeline
	//db.playlist.aggregate([{"$match": {"_id": ObjectId("6692d6e23a2f6303d4a3a331")}},{"$project": {"playlistname": 1,"createdAt": 1,"UpdatedAt": 1,"playedAt": 1,"isplaying": 1,"deviceblock": {"$filter": {"input":"$deviceblock","as": "block","cond": { "$eq": [ "$$block.deviceid", ObjectId("6692d6b83a2f6303d4a3a330") ] } } } } }]);
    pipeline := mongo.Pipeline{
        {{"$match", bson.D{{"_id", playlistobjectId}}}},
        {{"$project", bson.D{
            {"playlistname", 1},
            {"createdAt", 1},
            {"UpdatedAt", 1},
            {"playedAt", 1},
            {"isplaying", 1},
            {"deviceblock", bson.D{
                {"$filter", bson.D{
                    {"input", "$deviceblock"},
                    {"as", "block"},
                    {"cond", bson.D{
                        {"$eq", bson.A{"$$block.deviceid", screenobjectId}},
                    }},
                }},
            }},
        }}},
    }
    // Execute the aggregation
    cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return result, err
	}
    defer cursor.Close(ctx)

    // Iterate over the results
    err = cursor.All(ctx, &result)
	if err != nil {
		return result, err
	}
	
	return result, nil
}

func DeletePlaylist(ctx context.Context, client *mongo.Client, userID string, playlistID string) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")
	objectId, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("Number of documents deleted from Screens: %d\n", result.DeletedCount)
	return nil
}

func DuplicatePlaylist(ctx context.Context, client *mongo.Client, userID string, playlistID string) (DataModel.Playlist, error) {
	// Get playlst
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

	// duplicate playlist
	now := time.Now()
	result.CreatedAt = &now
	result.UpdatedAt = nil
	result.PlayedAt = nil
	result.Isplaying = false

	result.ID = primitive.NewObjectID()
	_, inserterr := coll.InsertOne(ctx, result)
	if inserterr != nil {
		return result, inserterr
	}
	log.Printf("Created User Playlist")
	return result, nil

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
	now := time.Now()
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
	coll := client.Database(userdsystemname).Collection("playlist")
	objectId, err := primitive.ObjectIDFromHex(playlistid)
	if err != nil {
		return err
	}
	stopcurrentplaylistfilter := bson.D{{"isplayling", true}}
	stopcurrentplaylistupdate := bson.D{{"$set",bson.D{{"isplaying",false}}}}
	stopresult, stopcurrentplaylistupdateerr := coll.UpdateOne(ctx, stopcurrentplaylistfilter, stopcurrentplaylistupdate)
	if stopcurrentplaylistupdateerr != nil {
		return stopcurrentplaylistupdateerr
	}
	fmt.Printf("Documents updated for stopping playlist: %v\n", stopresult.ModifiedCount)

	playcurrentplaylistfilter := bson.D{{"_id", objectId}}
	playcurrentplaylistupdate := bson.D{{"$set",bson.D{{"isplaying",true},{"playedAt",now}}}}
	playresult, playcuurentplaylisterr := coll.UpdateOne(ctx, playcurrentplaylistfilter, playcurrentplaylistupdate)
	if playcuurentplaylisterr != nil {
		return playcuurentplaylisterr
	}
	fmt.Printf("Documents updated for playlist playlist: %v\n", playresult.ModifiedCount)
	deviceIds,polerr := GetUniqueDeviceIds(ctx, client, userID, playlistid)
	if polerr != nil {
		return polerr
	}
	updatescreens := UpdateScreenCollection(ctx, client, userID, deviceIds, playlistid, playlist.Name)
	if updatescreens != nil{
		return updatescreens 
	}
	return nil

}

//Function to play playlist to a particular screen
func PlayPlaylisttoScreen(ctx context.Context, client *mongo.Client, userID string, playlistid string, screenid string) error {
	var playlist DataModel.Playlist
	_, err := GetUser(client, userID)
	if err != nil {
		return err
	}
	uservHostname := common.CreatevHostName(userID)
	userdsystemname := common.ExtractUserSystemIdentifier(userID)
	playlistarr, err := GetPlaylistwithSingleScreenData(ctx, client, userID, playlistid,screenid)

	if(len(playlistarr) == 1){
		playlist = playlistarr[0]
	} else {
		return errors.New("Playlist for the screen doesn't exist")
	}
	log.Printf("%v",playlist)
	rabbitqueue.Connect(userdsystemname, "password", uservHostname)
	err = rabbitqueue.PublishMessage(ctx, playlist, uservHostname)
	if err != nil {
		return err
	}

	screenobjectId, err := primitive.ObjectIDFromHex(screenid)
	deviceIds := []primitive.ObjectID{screenobjectId}
	updatescreens := UpdateScreenCollection(ctx, client, userID, deviceIds, playlistid, playlist.Name)
	if updatescreens != nil{
		return updatescreens 
	}
	return nil

}

// Function to get unique device IDs from a playlist
func GetUniqueDeviceIds(ctx context.Context, client *mongo.Client, userID string, playlistID string) ([]primitive.ObjectID, error) {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("playlist")

	playlistObjectID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", playlistObjectID}}}},
		{{"$unwind", "$deviceblock"}},
		{{"$group", bson.D{
			{"_id", nil},
			{"uniqueDeviceIds", bson.D{{"$addToSet", "$deviceblock.deviceid"}}},
		}}},
		{{"$project", bson.D{
			{"_id", 0},
			{"uniqueDeviceIds", 1},
		}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) > 0 {
		deviceIds := result[0]["uniqueDeviceIds"].(bson.A)
		uniqueDeviceIds := make([]primitive.ObjectID, len(deviceIds))
		for i, id := range deviceIds {
			uniqueDeviceIds[i] = id.(primitive.ObjectID)
		}
		return uniqueDeviceIds, nil
	}

	return nil, nil
}

// Function to update the screen collection
func UpdateScreenCollection(ctx context.Context, client *mongo.Client, userID string, deviceIDs []primitive.ObjectID, playlistID, playlistName string) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("screen")

	filter := bson.M{"_id": bson.M{"$in": deviceIDs}}
	update := bson.M{
		"$set": bson.M{
			"currentplaylistid":   playlistID,
			"currentplaylistname": playlistName,
		},
	}

	result, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	return nil
}