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
type PlayPlaylisttoScreenRequestjson struct {
	Playlistid string `bson:"playlistid"`
	Screenid string `bson:"screenidid"`
}
type PlayPlaylisttoScreenRequestjsonV2 struct {
	Playlistid []string `bson:"playlistid"`
	Screenid string `bson:"screenidid"`
}

func PlayPlaylistController(c *gin.Context) {
	var requestjsonvar PlayPlaylistRequestjson
	Userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
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

func PlayPlaylistControllerV2(c *gin.Context) {
	var requestjsonvar PlayPlaylistRequestjson
	Userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	err := dbmaster.PlayPlaylistV2(c, Client, Userid, requestjsonvar.Playlistid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successfully sent to Queue"})
	}
}
// PlayPlaylisttoScreen
func PlayPlaylisttoScreenController(c *gin.Context) {
	var requestjsonvar PlayPlaylisttoScreenRequestjson
	Userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	err := dbmaster.PlayPlaylisttoScreen(c, Client, Userid, requestjsonvar.Playlistid,requestjsonvar.Screenid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successfully sent to Queue"})
	}
}

func PlayPlaylisttoScreenControllerV2(c *gin.Context) {
	var requestjsonvar PlayPlaylisttoScreenRequestjsonV2
	Userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	err := dbmaster.PlayPlaylisttoScreenV2(c, Client, Userid, requestjsonvar.Playlistid,requestjsonvar.Screenid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successfully sent to Queue"})
	}
}
// Get playlist for a single screen data
func GetPlaylistofScreenController(c *gin.Context) {
	var requestjsonvar PlayPlaylisttoScreenRequestjson
	Userid := c.GetHeader("userid")
	if Userid == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No UserID Header Provided"})
	}

	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistofscreen, err := dbmaster.GetPlaylistwithSingleScreenData(c, Client, Userid, requestjsonvar.Playlistid,requestjsonvar.Screenid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": playlistofscreen})
	}
}



func CreatePlaylist(c *gin.Context) {
	var playlistjson DataModel.Playlist
	Userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
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
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentarray, err := dbmaster.ReadPlaylist(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetPlaylistbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	if userid  == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No UserID Header Provided"})
	}

	playlistId := c.Params.ByName("id")
	user, err := dbmaster.GetPlaylist(c, Client, userid, playlistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func DuplicatePlaylistbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	playlistId := c.Params.ByName("id")
	user, err := dbmaster.DuplicatePlaylist(c, Client, userid, playlistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func DeletePlaylistbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	playlistId := c.Params.ByName("id")
	err := dbmaster.DeletePlaylist(c, Client, userid, playlistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Successsfully Deleted PLaylist"})
	}
}

func UpdatePlaylistbyIDController(c *gin.Context) {
	var updatejson DataModel.UpdatePlaylistRequest
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
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


func AddScreenToPlaylistController(c *gin.Context) {
	type addscreentoplaylistrequestpayload struct{
		DefaultScreenID   string   `bson:"defaultscreenid,omitempty"`
		NewScreenID   string   `bson:"newscreenid,omitempty"`
		PlaylistID string 	`bson:"playlistid,omitempty"`
	}
	var addjson addscreentoplaylistrequestpayload
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&addjson)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	
	err := dbmaster.AddScreentoPlaylist(c, Client, userid, addjson.PlaylistID, addjson.DefaultScreenID,addjson.NewScreenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": "Successfully Updated"})
	}
}
