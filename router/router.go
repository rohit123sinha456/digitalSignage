package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rohit123sinha456/digitalSignage/common"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/objectstore"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var ObjectStoreClient *minio.Client
var R *gin.Engine

func SetupRouter() {
	R = gin.Default()
	Client = dbmaster.ConnectDB()
	ObjectStoreClient = objectstore.ConnectObjectStore()
}

func UserRouter() {
	R.GET("/user/", func(c *gin.Context) {
		users, err := dbmaster.GetAllUser(Client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"users": users})
		}
	})

	R.GET("/user/:id", func(c *gin.Context) {
		userID := c.Params.ByName("id")
		user, err := dbmaster.GetUser(Client, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	})

	R.POST("/user", func(c *gin.Context) {
		var newUser DataModel.User
		if c.Bind(&newUser) == nil {
			err := dbmaster.CreateUser(c, Client, ObjectStoreClient, newUser)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			}
			usid := common.ExtractUserSystemIdentifier(newUser.UserID)
			c.JSON(http.StatusOK, gin.H{"status": "ok", "userID": newUser.UserID, "userSystemID": usid})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Request is not in proper format"})
		}
	})
}
