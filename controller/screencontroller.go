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
	userid := c.GetHeader("userid")
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	screenid, err := dbmaster.CreateScreen(c, Client, userid, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"screenid": screenid})
	}
}

func ReadScreenController(c *gin.Context) {
	var contentarray []DataModel.Screen
	userid := c.GetHeader("userid")
	contentarray, err := dbmaster.ReadScreen(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetScreenbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
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
	userid := c.GetHeader("userid")
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
