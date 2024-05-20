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
	userid := c.GetHeader("userid")
	log.Printf("%+v", userid)
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	contentid, err := dbmaster.CreateContent(c, Client, userid, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"contentid": contentid})
	}
}

func ReadContentController(c *gin.Context) {
	var contentarray []DataModel.Content
	userid := c.GetHeader("userid")
	contentarray, err := dbmaster.ReadContent(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetContentbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	contentID := c.Params.ByName("id")
	user, err := dbmaster.ReadOneContent(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}
