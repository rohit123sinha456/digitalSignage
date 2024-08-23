package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"go.mongodb.org/mongo-driver/bson"
)

// type EventStreamRequest struct {
// 	Message string `form:"message" json:"message" binding:"required,max=100"`
// }
type EventStreamPostRequest struct {
	Screenmongoid string  `json:"screenmongoid"`
}

func CreateScreenController(c *gin.Context) {
	var requestjsonvar DataModel.Screen
	userid := c.GetHeader("userid")
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
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentarray, err := dbmaster.ReadScreen(c, Client, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"contents": contentarray})
}

func GetScreenbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	contentID := c.Params.ByName("id")
	user, err := dbmaster.ReadOneScreen(c, Client, userid, contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"content": user})
	}
}

func UpdateScreenbyIDController(c *gin.Context) {
	var requestjsonvar DataModel.Screen
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	screenID := c.Params.ByName("id")
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
		return
	}
	err := dbmaster.UpdateScreen(c, Client, userid, screenID, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Updated"})
	}
}

func PublicUpdateScreenbyIDController(c *gin.Context) {
	var requestjsonvar DataModel.Screen
	screenID := c.Params.ByName("id")
	userid := c.Params.ByName("userid")
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
		return
	}
	err := dbmaster.UpdateScreen(c, Client, userid, screenID, requestjsonvar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Updated"})
	}
}


func DeleteScreenbyIDController(c *gin.Context) {
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	screenID := c.Params.ByName("id")
	err := dbmaster.DeleteScreen(c, Client, userid, screenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
	}
}


func HandleEventStreamPost(c *gin.Context, ch chan DataModel.EventStreamRequest, screencode string) {
	var result DataModel.EventStreamRequest
	var userinfo DataModel.UserSystemIdentifeir
	var requestbody EventStreamPostRequest

	coll := Client.Database("user").Collection("userSystemInfo")
	userid := c.GetHeader("userid")
	value, ifexists := c.Get("uid")
	if ifexists == true {
		log.Printf("%s", value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User Id In Token"})
	}
	reqerr := c.Bind(&requestbody)
	log.Printf("%+v", requestbody)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}

	filter := bson.D{{"userid", userid}}
	err := coll.FindOne(c, filter).Decode(&userinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	result.Screencode = screencode
	result.Userinfo = userinfo
	result.ScreenMongoID = requestbody.Screenmongoid

	log.Printf("%+v", result)
	// if err := c.Bind(&result); err != nil {
	// 	log.Fatal(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"Bind Status": err.Error()})
	// 	return
	// }
	ch <- result
	c.JSON(http.StatusOK, gin.H{"status": "Updated"})
	return
}

func HandleEventStreamGet(c *gin.Context, ch chan DataModel.EventStreamRequest) {
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-ch; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})

	return
}
