package router

import (
	"log"
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

type AddDeviceRequestjson struct {
	Userid  string             `bson:"userid"`
	Devices []DataModel.Device `bson:"devices"`
}

type PlayPlaylistRequestjson struct {
	Action     DataModel.PlaylistActionType `bson:"action"`
	Userid     string                       `bson:"userid"`
	Playlistid string                       `bson:"playlistid"`
}

type CreatePlaylistRequestjson struct {
	Userid   string             `bson:"userid"`
	Playlist DataModel.Playlist `bson:"playlist"`
}

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

func DeviceRouter() {
	R.POST("/device", func(c *gin.Context) {
		var requestjsonvar AddDeviceRequestjson
		reqerr := c.Bind(&requestjsonvar)
		log.Printf("%+v", requestjsonvar)
		if reqerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
		}
		userID, err := dbmaster.CreateDevice(c, Client, requestjsonvar.Userid, requestjsonvar.Devices)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		}
		user, err := dbmaster.GetUser(Client, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user})
		}

	})
}

func PlaylistRouter() {
	R.POST("/playplaylist", func(c *gin.Context) {
		var requestjsonvar PlayPlaylistRequestjson
		reqerr := c.Bind(&requestjsonvar)
		log.Printf("%+v", requestjsonvar)
		if reqerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
		}
		err := dbmaster.PlayPlaylist(c, Client, requestjsonvar.Userid, requestjsonvar.Playlistid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Successfully sent to Queue"})
		}
	})
	R.POST("/playlist", func(c *gin.Context) {
		var requestjsonvar CreatePlaylistRequestjson
		reqerr := c.Bind(&requestjsonvar)
		log.Printf("%+v", requestjsonvar)
		if reqerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
		}
		playlistid, err := dbmaster.CreatePlaylist(c, Client, requestjsonvar.Userid, requestjsonvar.Playlist)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
		}
	})
}
