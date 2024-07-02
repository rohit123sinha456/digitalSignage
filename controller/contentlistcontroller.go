package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

func CreateContentListController(c *gin.Context) {
	var requestjsonvar DataModel.ContentList
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	log.Printf("Content Block")
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistid, err := dbmaster.CreateContentList(c, Client, userid, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
	}
}

func ReadContentListController(c *gin.Context) {
	var contentarray []DataModel.ContentList
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentarray, err := dbmaster.ReadContentList(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}
