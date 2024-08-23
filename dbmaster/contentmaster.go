package dbmaster

import (
	"time"
	"context"
	"log"
	"net/url"
	"strings"
	"mime/multipart"
	"github.com/rohit123sinha456/digitalSignage/common"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/minio/minio-go/v7"
	"github.com/rohit123sinha456/digitalSignage/config"
	"github.com/rohit123sinha456/digitalSignage/objectstore"

)

func CreateContent(ctx context.Context, client *mongo.Client, userID string, content DataModel.Content) (string, error) {
	now := time.Now()
	content.CreatedAt = &now
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

func UpdateContent(ctx context.Context, client *mongo.Client, userID string, contentID string, updateRequest DataModel.Content) error {
	userSystemname := common.ExtractUserSystemIdentifier(userID)
	coll := client.Database(userSystemname).Collection("content")
	objectId, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	 // Build the update document
	 updateFields := bson.D{}
	 if updateRequest.CName != "" {
		 updateFields = append(updateFields, bson.E{"cname", updateRequest.CName})
	 }
	 if updateRequest.DType != "" {
		 updateFields = append(updateFields, bson.E{"dtype", updateRequest.DType})
	 }
	 if updateRequest.Link != "" {
		 updateFields = append(updateFields, bson.E{"link", updateRequest.Link})
	 }
	 if updateRequest.CreatedAt != nil {
		 updateFields = append(updateFields, bson.E{"createdAt", updateRequest.CreatedAt})
	 }

	 if len(updateFields) == 0 {
		 return nil
	 }
	
	update := bson.D{{"$set", updateFields}}
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Printf("Documents updated: %v\n", result.ModifiedCount)
	return nil
}

func UploadContent(ctx context.Context, objectStoreClient *minio.Client, userID string,filedata *multipart.FileHeader) (string,error) {
	userBucketname := common.CreateBucketName(userID)
	log.Printf("Uploading Content")
	uploaderr := objectstore.StoreFile(ctx, objectStoreClient, userBucketname,filedata)
	if uploaderr != nil {
		return "",uploaderr
	}
	sourceurl := config.GetEnvbyKey("OBJECTSTOREURL")
	nospaces := strings.TrimSpace(filedata.Filename)
	validfile := url.PathEscape(nospaces)
	objecturl := sourceurl + userBucketname + "/"+ validfile
	log.Printf("Successfull Content Uploaded contentmaster")
	return objecturl,nil
}

// func UpdateContent() {}
