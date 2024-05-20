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

func UserRouter() { //Done
	// R.GET("/user/", controller.GetAllUserController)

	// R.GET("/user/:id", controller.GetUserbyIDController)

	// R.POST("/user", controller.CreateNewUserController)
	R.POST("users/signup", controller.Signup)
	R.POST("users/login", controller.Login)
}

func AuthRoutes() {
	// R.Use(middleware.Authenticate())
	R.GET("/usersdata", controller.GetAllUserController)
}

func PlaylistRouter() { // Done
	R.POST("/playplaylist", controller.PlayPlaylistController)
	R.POST("/playlist", controller.CreatePlaylist)                   // Create Playlist
	R.GET("/playlist", controller.ReadPlaylistController)            // Read (all)
	R.GET("/playlist/:id", controller.GetPlaylistbyIDController)     // Read (Specific)
	R.POST("/playlist/:id", controller.UpdatePlaylistbyIDController) // Update (Specific)

}

func ContentRouter() { // Done
	R.POST("/content", controller.CreateContentController)     // create Content
	R.GET("/content", controller.ReadContentController)        // Read (all)
	R.GET("/content/:id", controller.GetContentbyIDController) // Read (Specific)

}

func ContentListRouter() {
	R.POST("/contentlist", controller.CreateContentListController)
	// R.GET("/contentlist", controller.ReadContentController)
	// R.GET("/contentlist/:id", controller.GetContentbyIDController)

}

func ScreenRouter() { //Done
	R.POST("/screen", controller.CreateScreenController)
	R.GET("/screen", controller.ReadScreenController)
	R.GET("/screen/:id", controller.GetScreenbyIDController)
	R.POST("/screen/:id", controller.UpdateScreenbyIDController) // Not working

}
