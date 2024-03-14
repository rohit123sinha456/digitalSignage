package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

func CreateScreenController(c *gin.Context) {
	var requestjsonvar DataModel.Screen
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistid, err := dbmaster.CreateScreen(c, Client, userid, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
	}
}

func ReadScreenController(c *gin.Context) {
	var contentarray []DataModel.Screen
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	contentarray, err := dbmaster.ReadScreen(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetScreenbyIDController(c *gin.Context) {
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	contentID := c.Params.ByName("id")
	user, err := dbmaster.ReadOneScreen(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func UpdateScreenbyIDController(c *gin.Context) {
	var requestjsonvar []DataModel.ScreenBlock
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	screenID := c.Params.ByName("id")
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	err := dbmaster.UpdateScreen(c, Client, userid, screenID, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Updated"})
	}
}
