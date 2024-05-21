package dbmaster

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/rohit123sinha456/digitalSignage/common"
	"github.com/rohit123sinha456/digitalSignage/config"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/objectstore"
	"github.com/rohit123sinha456/digitalSignage/rabbitqueue"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func ConnectDB() *mongo.Client {
	uri := config.GetEnvbyKey("APPDB")//os.Getenv("APPDB") //"mongodb://localhost:27017"
	fmt.Println(uri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if client != nil {
		log.Printf("Client successfully Created")
	}
	log.Printf("Successfully Connected to database")
	return client
}

func CreateUserSystemInfo(ctx context.Context, client *mongo.Client, newUser DataModel.User) error {
	var usersystemidentifier DataModel.UserSystemIdentifeir
	userID := newUser.UserID
	usersystemidentifier.UserID = userID
	usersystemidentifier.UserQueuevHostID = common.CreatevHostName(userID)
	usersystemidentifier.UserSystemID = common.ExtractUserSystemIdentifier(userID)
	usersystemidentifier.UserQueueID = common.ExtractUserSystemIdentifier(userID)
	usersystemidentifier.UserBucketID = common.CreateBucketName(userID)

	collection := client.Database("user").Collection("userSystemInfo")
	_, err := collection.InsertOne(ctx, usersystemidentifier)
	if err != nil {
		return err
	}
	return nil
}

func GetUserSystemInfo(ctx context.Context, client *mongo.Client, userId string) (DataModel.UserSystemIdentifeir, error) {
	var result DataModel.UserSystemIdentifeir
	coll := client.Database("user").Collection("userSystemInfo")
	usersystemid := common.ExtractUserSystemIdentifier(userId)
	filter := bson.D{{"usersystemid", usersystemid}}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Create a transactional atomic property that if half way through the systen fails, all the previous operations
// gets reverted. i.e if record entered in db and bucket creaed but rabbitmq failed, the users needs tto be deleted
// from the DB and the bucket needs to be detected as well
func CreateUser(ctx context.Context, client *mongo.Client, objectStoreClient *minio.Client, newUser DataModel.User) (string, error) {
	userID := uuid.NewString()
	doesUserExists, _ := GetUser(client, userID)
	doesUserSystemInfoExists, _ := GetUserSystemInfo(ctx, client, userID)
	if reflect.ValueOf(doesUserExists).IsZero() != true || reflect.ValueOf(doesUserSystemInfoExists).IsZero() != true {
		return "", errors.New(" Generated UserId Already Exists, Please try again")
	}
	newUser.UserID = userID
	coll := client.Database("user").Collection("userData")
	_, err := coll.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}
	log.Printf("User Successfully inserted in Database")
	uservhostname := common.CreatevHostName(userID)
	userdsystemname := common.ExtractUserSystemIdentifier(userID)
	userbucketname := common.CreateBucketName(userID)

	obserror := objectstore.CreateBucket(ctx, objectStoreClient, userbucketname)
	queueerr := rabbitqueue.SetupUserandvHost(userdsystemname, uservhostname)
	metainfoerr := CreateUserSystemInfo(ctx, client, newUser)
	log.Printf("Created User")
	if queueerr != nil {
		return "", queueerr
	} else if obserror != nil {
		return "", obserror
	} else if metainfoerr != nil {
		return "", metainfoerr
	} else {
		return userID, nil
	}
}

// Modify the function so that base on the string length, it uses fileds of systemID, vHostID and so on from userSystemInfo
// fetch the userID and then fetch the result from userData
func GetUser(client *mongo.Client, userId string) (DataModel.User, error) {
	var result DataModel.User
	coll := client.Database("user").Collection("userData")
	filter := bson.D{{"userid", userId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllUser(client *mongo.Client) ([]DataModel.User, error) {
	var results []DataModel.User
	coll := client.Database("user").Collection("userData")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return results, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}

func AddUserDevice(client *mongo.Client, userId string) {
	var result DataModel.User
	coll := client.Database("user").Collection("userData")
	filter := bson.D{{"user_id", userId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
