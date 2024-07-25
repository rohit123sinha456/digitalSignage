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
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
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
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentarray, err := dbmaster.ReadContent(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetContentbyIDController(c *gin.Context) {
	contentID := c.Params.ByName("id")
	userid := c.Params.ByName("userid")
	user, err := dbmaster.ReadOneContent(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func DeleteContentbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentID := c.Params.ByName("id")
	err := dbmaster.DeleteContent(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
	}
}

func UploadContentController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		log.Printf("%s", value)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	file, err := c.FormFile("fileUpload")

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"statuses": err.Error()})
		return
    }

	objecturl,uploaderr := dbmaster.UploadContent(c,ObjectStoreClient,userid,file)
	if uploaderr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
    }
	c.JSON(http.StatusOK, gin.H{"URL": objecturl})
}