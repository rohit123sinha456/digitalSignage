package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rohit123sinha456/digitalSignage/common"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	helper "github.com/rohit123sinha456/digitalSignage/helper"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var Client *mongo.Client
var ObjectStoreClient *minio.Client

func SetupUserController(mongoclient *mongo.Client, obsclient *minio.Client) {
	Client = mongoclient
	ObjectStoreClient = obsclient
}

func GetAllUserController(c *gin.Context) {
	users, err := dbmaster.GetAllUser(Client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func GetUserbyIDController(c *gin.Context) {
	userID := c.Params.ByName("id")
	user, err := dbmaster.GetUser(Client, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func CreateNewUserController(c *gin.Context) {
	var newUser DataModel.User
	if c.Bind(&newUser) == nil {
		_, err := dbmaster.CreateUser(c, Client, ObjectStoreClient, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			usid := common.ExtractUserSystemIdentifier(newUser.UserID)
			c.JSON(http.StatusOK, gin.H{"status": "ok", "userID": newUser.UserID, "userSystemID": usid})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Request is not in proper format"})
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("E-Mail or Password is incorrect")
		check = false
	}
	return check, msg
}

func Signup(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user DataModel.User
	userCollection := Client.Database("user").Collection("userData")
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		// log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error detected while fetching the email"})
		return
	}

	password := HashPassword(*user.Password)
	user.Password = &password

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		// log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occured while fetching the phone number"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "The mentioned E-Mail or Phone Number already exists"})
		return
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	// userappid, insertErr := dbmaster.CreateUser(c, Client, ObjectStoreClient, user)
	log.Printf("Starting Signup")
	userappid, insertErr := dbmaster.TransactionCreateUser(c, Client, ObjectStoreClient, user)
	if insertErr != nil {
		msg := fmt.Sprintf("User Details were not Saved")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	inserteduser, err := dbmaster.GetUser(Client, userappid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": inserteduser})
	}
	// c.JSON(http.StatusOK, userappid)

}



func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user DataModel.User
	var foundUser DataModel.User
	userCollection := Client.Database("user").Collection("userData")

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	if foundUser.Email == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
	}
	token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.UserID)
	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foundUser)

}


func Logout(c * gin.Context){
	coll := Client.Database("user").Collection("userData")
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	filter := bson.D{{"userid", userid}}
	update := bson.D{{"$set", bson.D{{"token", "NoStringsAttached"},{"refresh_token","NoRefreshToken"}}}}
	_, err := coll.UpdateOne(c, filter, update)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "User Logged Out")
}