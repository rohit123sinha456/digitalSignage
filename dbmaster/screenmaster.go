package dbmaster

import (
	"context"
	"errors"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

)
type  ActivatePlaylistofScreen struct{
	CurrentPlaylistName string  `bson:"playlistname,omitempty"`
	CurrentPlaylistID primitive.ObjectID    `bson:"_id,omitempty"`
}

func checkifcontentlistexists(ctx context.Context, coll *mongo.Collection, screenblocks []DataModel.ScreenBlock) error {
	log.Printf("Checking Contents")
	var objectIDs []primitive.ObjectID

	for _, screenblock := range screenblocks {
		log.Printf("%s", screenblock.BlockName)
		contentlistId := screenblock.ContentListID
		objectIDs = append(objectIDs, contentlistId)
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

func CreateScreen(ctx context.Context, client *mongo.Client, userID string, screendetails DataModel.Screen) (string, error) {
	now := time.Now()
	screendetails.CreatedAt = &now
	screendetails.ID = primitive.NewObjectID()
	userdBname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userdBname).Collection("screen")
	result, inserterr := coll.InsertOne(ctx, screendetails)
	if inserterr != nil {
		return "", inserterr
	}
	log.Printf("Screen Created")
	idstring := result.InsertedID.(primitive.ObjectID).Hex()
	return idstring, nil
}

func ReadScreen(ctx context.Context, client *mongo.Client, userID string) ([]DataModel.Screen, error) {
	var contentlistarray []DataModel.Screen
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("screen")
	filter := bson.D{{}}
	curr, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		var result DataModel.Screen
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}
		contentlistarray = append(contentlistarray, result)
	}
	defer curr.Close(ctx)
	return contentlistarray, nil
}

func ReadOneScreen(ctx context.Context, client *mongo.Client, userID string, screenID string) (DataModel.Screen, error) {
	var result DataModel.Screen
	// var activateplaylistofscreen []ActivatePlaylistofScreen
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("screen")
	objectId, err := primitive.ObjectIDFromHex(screenID)
	if err != nil {
		return result, err
	}
	filter := bson.D{{"_id", objectId}}
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//db.playlist.find({$and:[{isplaying:true},{deviceblock:{$elemMatch:{deviceid:ObjectId('6692d6b83a2f6303d4a3a330')}}}]},{_id:1,playlistname:1})
	// currentplaylistfilter := bson.D{
	// 	{"$and", bson.A{
	// 		bson.D{{"isplaying", true}},
	// 		bson.D{{"deviceblock", bson.D{{"$elemMatch", bson.D{{"deviceid", objectId}}}}}},
	// 	}},
	// }

	// currentplaylistprojection := bson.D{
	// 	{"_id", 1},
	// 	{"playlistname", 1},
	// }

	// opts := options.Find().SetProjection(currentplaylistprojection)
	// playlistcoll := client.Database(userSystemname).Collection("playlist")
	// cursor, err := playlistcoll.Find(context.TODO(), currentplaylistfilter, opts)
	// if err != nil {
	// 	panic(err)
	// }
	// if err = cursor.All(context.TODO(), &activateplaylistofscreen); err != nil {
	// 	panic(err)
	// }
	// log.Printf("Feting User details")
	// log.Printf("%v",activateplaylistofscreen)
	// result.CurrentPlaylistName = activateplaylistofscreen[0].CurrentPlaylistName
	// result.CurrentPlaylistID = activateplaylistofscreen[0].CurrentPlaylistID
	return result, nil
}

func UpdateScreen(ctx context.Context, client *mongo.Client, userID string, screenID string, updateRequest DataModel.Screen) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("screen")
	objectId, err := primitive.ObjectIDFromHex(screenID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	// Build the update document
	updateFields := bson.D{}
	if updateRequest.Name != "" {
		updateFields = append(updateFields, bson.E{"name", updateRequest.Name})
	}
	if updateRequest.Location != "" {
		updateFields = append(updateFields, bson.E{"location", updateRequest.Location})
	}
	if updateRequest.CurrentPlaylistName != "" {
		updateFields = append(updateFields, bson.E{"currentplaylistname", updateRequest.CurrentPlaylistName})
	}
	if updateRequest.CurrentPlaylistID != primitive.NilObjectID  {
		updateFields = append(updateFields, bson.E{"currentplaylistid", updateRequest.CurrentPlaylistID})
	}
	if updateRequest.CreatedAt != nil {
		updateFields = append(updateFields, bson.E{"createdAt", updateRequest.CreatedAt})
	}
	if updateRequest.UpdatedAt != nil {
		updateFields = append(updateFields, bson.E{"updatedAt", updateRequest.UpdatedAt})
	}
	if updateRequest.Status {
		updateFields = append(updateFields, bson.E{"status", updateRequest.Status})
	}
	if updateRequest.Orientation != 0 {
		updateFields = append(updateFields, bson.E{"orientation", updateRequest.Orientation})
	}
	if updateRequest.StorageTotal != 0 {
		updateFields = append(updateFields, bson.E{"storagetotal", updateRequest.StorageTotal})
	}
	if updateRequest.StorageFree != 0 {
		updateFields = append(updateFields, bson.E{"storagefree", updateRequest.StorageFree})
	}
	if updateRequest.StorageUsed != 0 {
		updateFields = append(updateFields, bson.E{"storageused", updateRequest.StorageUsed})
	}
	if updateRequest.MemoryTotal != 0 {
		updateFields = append(updateFields, bson.E{"memorytotal", updateRequest.MemoryTotal})
	}
	if updateRequest.MemoryUsed != 0 {
		updateFields = append(updateFields, bson.E{"memoryused", updateRequest.MemoryUsed})
	}
	if updateRequest.IPAddr != "" {
		updateFields = append(updateFields, bson.E{"ip", updateRequest.IPAddr})
	}
	if updateRequest.DeviceModel != "" {
		updateFields = append(updateFields, bson.E{"devicemodel", updateRequest.DeviceModel})
	}
	if updateRequest.CanDrawOverlay {
		updateFields = append(updateFields, bson.E{"candrawoverlay", updateRequest.CanDrawOverlay})
	}
	if updateRequest.AppShellVersion != "" {
		updateFields = append(updateFields, bson.E{"appshellversion", updateRequest.AppShellVersion})
	}
	if updateRequest.ScreenshotSupport {
		updateFields = append(updateFields, bson.E{"screenshotsupport", updateRequest.ScreenshotSupport})
	}
	if updateRequest.ScreenResolution != "" {
		updateFields = append(updateFields, bson.E{"screenresolution", updateRequest.ScreenResolution})
	}
	if updateRequest.BrowserResolution != "" {
		updateFields = append(updateFields, bson.E{"browserresolution", updateRequest.BrowserResolution})
	}
	if updateRequest.EngerySavedEnabled {
		updateFields = append(updateFields, bson.E{"energysaver", updateRequest.EngerySavedEnabled})
	}
	if updateRequest.Country != "" {
		updateFields = append(updateFields, bson.E{"country", updateRequest.Country})
	}
	if updateRequest.UserAgent != "" {
		updateFields = append(updateFields, bson.E{"useragent", updateRequest.UserAgent})
	}
	if updateRequest.GPlaySupport {
		updateFields = append(updateFields, bson.E{"gplaysupport", updateRequest.GPlaySupport})
	}
	if updateRequest.VideoCodecs != "" {
		updateFields = append(updateFields, bson.E{"videocodecs", updateRequest.VideoCodecs})
	}
	if updateRequest.PlayerTimezone != "" {
		updateFields = append(updateFields, bson.E{"playertimsezone", updateRequest.PlayerTimezone})
	}
	if updateRequest.OS != "" {
		updateFields = append(updateFields, bson.E{"os", updateRequest.OS})
	}
	if updateRequest.DevicePixelRatio != "" {
		updateFields = append(updateFields, bson.E{"devicepixelratio", updateRequest.DevicePixelRatio})
	}
	if updateRequest.PlayerCodec != "" {
		updateFields = append(updateFields, bson.E{"playercodec", updateRequest.PlayerCodec})
	}
	if updateRequest.RicoviAppVersion != "" {
		updateFields = append(updateFields, bson.E{"appversion", updateRequest.RicoviAppVersion})
	}
	update := bson.D{{"$set", updateFields}}

	// Updates the first document that has the specified "_id" value
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Printf("Documents updated: %v\n", result.ModifiedCount)
	return nil
}

func DeleteScreen(ctx context.Context, client *mongo.Client, userID string, screenID string) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	screenCollection := client.Database(userSystemname).Collection("screen")
	userdBname := common.ExtractUserSystemIdentifier(userID)
	plalistCollection := client.Database(userdBname).Collection("playlist")

	objectId, err := primitive.ObjectIDFromHex(screenID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	result, err := screenCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("Number of documents deleted from Screens: %d\n", result.DeletedCount)
	// db.playlist.updateMany({deviceblock:{$elemMatch:{deviceid:ObjectId('6683a8e6a0cf1a28f6edddd7')}}},{$pull:{deviceblock:{deviceid: ObjectId('6683a8e6a0cf1a28f6edddd7')}}})
	playlistfilter := bson.M{
        "deviceblock": bson.M{
            "$elemMatch": bson.M{
                "deviceid": objectId,
            },
        },
    }
    playlistupdate := bson.M{
        "$pull": bson.M{
            "deviceblock": bson.M{
                "deviceid": objectId,
            },
        },
    }

    // Perform the update operation
    playlistdeleteresult, err := plalistCollection.UpdateMany(ctx, playlistfilter, playlistupdate)
    if err != nil {
        return err
    }
	log.Printf("Number of Screens deleted from Plalist: %d\n", playlistdeleteresult.ModifiedCount)
	return nil
}