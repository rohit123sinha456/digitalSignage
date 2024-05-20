package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

type PlayPlaylistRequestjson struct {
	Playlistid string `bson:"playlistid"`
}

func PlayPlaylistController(c *gin.Context) {
	var requestjsonvar PlayPlaylistRequestjson
	Userid := c.GetHeader("userid")

	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	err := dbmaster.PlayPlaylist(c, Client, Userid, requestjsonvar.Playlistid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successfully sent to Queue"})
	}
}

func CreatePlaylist(c *gin.Context) {
	var playlistjson DataModel.Playlist
	Userid := c.GetHeader("userid")
	reqerr := c.Bind(&playlistjson)
	log.Printf("%+v", playlistjson)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistid, err := dbmaster.CreatePlaylist(c, Client, Userid, playlistjson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
	}
}

func ReadPlaylistController(c *gin.Context) {
	var contentarray []DataModel.Playlist
	userid := c.GetHeader("userid")
	contentarray, err := dbmaster.ReadPlaylist(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetPlaylistbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	playlistId := c.Params.ByName("id")
	user, err := dbmaster.GetPlaylist(c, Client, userid, playlistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func UpdatePlaylistbyIDController(c *gin.Context) {
	var updatejson DataModel.UpdatePlaylistRequest
	userid := c.GetHeader("userid")
	reqerr := c.Bind(&updatejson)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistId := c.Params.ByName("id")
	err := dbmaster.UpdatePlaylist(c, Client, userid, playlistId, updatejson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": "Successfully Updated"})
	}
}
