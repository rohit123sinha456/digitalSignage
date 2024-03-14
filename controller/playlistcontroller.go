package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
)

type PlayPlaylistRequestjson struct {
	Action     string `bson:"action"`
	Userid     string `bson:"userid"`
	Playlistid string `bson:"playlistid"`
}

type CreatePlaylistRequestjson struct {
	Screenid string `bson:"screenid"`
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
	Userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistid, err := dbmaster.CreatePlaylist(c, Client, Userid, requestjsonvar.Screenid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
	}
}
