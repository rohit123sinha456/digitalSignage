package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

func CreateContentController(c *gin.Context) {
	var requestjsonvar DataModel.Content
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	playlistid, err := dbmaster.CreateContent(c, Client, userid, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"playlistid": playlistid})
	}
}

func ReadContentController(c *gin.Context) {
	var contentarray []DataModel.Content
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	contentarray, err := dbmaster.ReadContent(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetContentbyIDController(c *gin.Context) {
	userid := "dd50c75c-7509-4f66-b312-a98445c6c65c"
	contentID := c.Params.ByName("id")
	user, err := dbmaster.ReadOneContent(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}
