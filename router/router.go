package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rohit123sinha456/digitalSignage/controller"
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
	"github.com/rohit123sinha456/digitalSignage/middleware"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
	"github.com/rohit123sinha456/digitalSignage/objectstore"
)

var R *gin.Engine
var public *gin.RouterGroup
var private *gin.RouterGroup

func SetupRouter() {
	R = gin.Default()
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	config.AllowAllOrigins = true
	// R.Use(cors.Default())
	R.Use(cors.New(config))
	private = R.Group("/api")
	private.Use(middleware.Authenticate())
	public = R.Group("/api/public")
	Client := dbmaster.ConnectDB()
	ObjectStoreClient := objectstore.ConnectObjectStore()
	controller.SetupUserController(Client, ObjectStoreClient)
}

func UserRouter() { //Done
	// R.GET("/user/", controller.GetAllUserController)

	// R.GET("/user/:id", controller.GetUserbyIDController)

	// R.POST("/user", controller.CreateNewUserController)
	public.POST("users/signup", controller.Signup) // This will go to Admin Section
	public.POST("users/login", controller.Login)
}

func AuthRoutes() {
	// R.Use(middleware.Authenticate())
	private.GET("/usersdata", controller.GetAllUserController)
}

func PlaylistRouter() { // Done
	private.POST("/playplaylist", controller.PlayPlaylistController)
	private.POST("/playlist", controller.CreatePlaylist)                   // Create Playlist
	private.GET("/playlist", controller.ReadPlaylistController)            // Read (all)
	private.GET("/playlist/:id", controller.GetPlaylistbyIDController)     // Read (Specific)
	private.POST("/playlist/:id", controller.UpdatePlaylistbyIDController) // Update (Specific)

}

func ContentRouter() { // Done
	private.POST("/content", controller.CreateContentController)     // create Content
	private.GET("/content", controller.ReadContentController)        // Read (all)
	private.GET("/content/:id", controller.GetContentbyIDController) // Read (Specific)

}

func ContentListRouter() {
	private.POST("/contentlist", controller.CreateContentListController)
	// R.GET("/contentlist", controller.ReadContentController)
	// R.GET("/contentlist/:id", controller.GetContentbyIDController)

}

func ScreenRouter() { //Done
	ch := make(chan DataModel.EventStreamRequest)
	private.POST("/screen", controller.CreateScreenController)
	private.GET("/screen", controller.ReadScreenController)
	private.GET("/screen/:id", controller.GetScreenbyIDController)
	private.POST("/screen/:id", controller.UpdateScreenbyIDController) // Not working

	private.POST("/event-stream/:id", func(c *gin.Context) {
		screencode := c.Params.ByName("id")
		controller.HandleEventStreamPost(c, ch, screencode)
	})
	public.GET("/event-stream/:id", middleware.HeadersMiddleware(), func(c *gin.Context) {
		controller.HandleEventStreamGet(c, ch)
	})
}
