package controller


import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"

)

type AddDeviceRequestjson struct {
	Userid  string             `bson:"userid"`
	Devices []DataModel.Device `bson:"devices"`
}


func CreateNewDeviceController(c *gin.Context) {
	var requestjsonvar AddDeviceRequestjson
	reqerr := c.Bind(&requestjsonvar)
	log.Printf("%+v", requestjsonvar)
	if reqerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": reqerr.Error()})
	}
	userID, err := dbmaster.CreateDevice(c, Client, requestjsonvar.Userid, requestjsonvar.Devices)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	user, err := dbmaster.GetUser(Client, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}

}
