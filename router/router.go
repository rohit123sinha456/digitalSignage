package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/controller"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	"github.com/rohit123sinha456/digitalSignage/objectstore"
)

var R *gin.Engine

func SetupRouter() {
	R = gin.Default()
	Client := dbmaster.ConnectDB()
	ObjectStoreClient := objectstore.ConnectObjectStore()
	controller.SetupUserController(Client, ObjectStoreClient)
}

func UserRouter() {
	R.GET("/user/", controller.GetAllUserController)

	R.GET("/user/:id", controller.GetUserbyIDController)

	R.POST("/user", controller.CreateNewUserController)
}

func DeviceRouter() {
	R.POST("/device", controller.CreateNewDeviceController)
}

func PlaylistRouter() {
	R.POST("/playplaylist", controller.PlayPlaylistController)
	R.POST("/playlist", controller.CreatePlaylist)
}
