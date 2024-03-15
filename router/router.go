package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/controller"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	"github.com/rohit123sinha456/digitalSignage/middleware"
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
	R.POST("users/signup", controller.Signup)
	R.POST("users/login", controller.Login)
}

func AuthRoutes() {
	R.Use(middleware.Authenticate())
	R.GET("/usersdata", controller.GetAllUserController)
}

func DeviceRouter() {
	R.POST("/device", controller.CreateNewDeviceController)
}

func PlaylistRouter() {
	R.POST("/playplaylist", controller.PlayPlaylistController)
	R.POST("/playlist", controller.CreatePlaylist)
}

func ContentRouter() {
	R.POST("/content", controller.CreateContentController)
	R.GET("/content", controller.ReadContentController)
	R.GET("/content/:id", controller.GetContentbyIDController)

}

func ContentListRouter() {
	R.POST("/contentlist", controller.CreateContentListController)
	// R.GET("/contentlist", controller.ReadContentController)
	// R.GET("/contentlist/:id", controller.GetContentbyIDController)

}

func ScreenRouter() {
	R.POST("/screen", controller.CreateScreenController)
	R.GET("/screen", controller.ReadScreenController)
	R.GET("/screen/:id", controller.GetScreenbyIDController)
	R.POST("/screen/:id", controller.UpdateScreenbyIDController)

}
