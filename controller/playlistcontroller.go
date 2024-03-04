package controller


import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"

)

type PlayPlaylistRequestjson struct {
	Action     DataModel.PlaylistActionType `bson:"action"`
	Userid     string                       `bson:"userid"`
	Playlistid string                       `bson:"playlistid"`
}

type CreatePlaylistRequestjson struct {
	Userid   string             `bson:"userid"`
	Playlist DataModel.Playlist `bson:"playlist"`
}

func PlayPlaylistController(c *gin.Context) {
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
}

func CreatePlaylist(c *gin.Context) {
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
}
