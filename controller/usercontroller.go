package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rohit123sinha456/digitalSignage/common"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	"go.mongodb.org/mongo-driver/mongo"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"

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
		err := dbmaster.CreateUser(c, Client, ObjectStoreClient, newUser)
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
