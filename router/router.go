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
	R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},}))
	private = R.Group("/api")
	private.Use(middleware.Authenticate())
	// private.Use(middleware.CORSMiddleware())
	public = R.Group("/api/public")
	// public.Use(middleware.CORSMiddleware())
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
	public.POST("users/getotp", controller.GenerateOTPController)
	public.POST("users/resetpassword", controller.ResetPassword)
	private.POST("users/logout", controller.Logout)

}

func AuthRoutes() {
	// R.Use(middleware.Authenticate())
	private.GET("/usersdata", controller.GetAllUserController)
}

func PlaylistRouter() { // Done
	private.POST("/playplaylist", controller.PlayPlaylistController)		// PLay PLaylist to all registered screen
	private.POST("/playplaylisttoscreen", controller.PlayPlaylisttoScreenController) // PLay PLaylist to single registered screen
	
	private.POST("/v2/playplaylist", controller.PlayPlaylistControllerV2)		// PLay PLaylist to all registered screen
	private.POST("/v2/playplaylisttoscreen", controller.PlayPlaylisttoScreenControllerV2) // PLay PLaylist to single registered screen
	
	private.POST("/playlist", controller.CreatePlaylist)                   // Create Playlist
	private.GET("/playlist", controller.ReadPlaylistController)            // Read (all)
	private.GET("/playlist/:id", controller.GetPlaylistbyIDController)     // Read (Specific)
	public.GET("/playlist/:id", controller.GetPlaylistbyIDController)     // APP - Read (Specific)
	private.POST("/playlist/:id", controller.UpdatePlaylistbyIDController) // Update (Specific)
	private.POST("/playlist/addscreen", controller.AddScreenToPlaylistController) // Add Scrren to Plalist (Specific)

	private.POST("/playlist/duplicate/:id", controller.DuplicatePlaylistbyIDController) // Duplicate (Specific ID)
	private.DELETE("/playlist/:id", controller.DeletePlaylistbyIDController) // Delete (Specific)
	

}

func ContentRouter() { // Done
	private.POST("/content", controller.CreateContentController)     // create Content
	private.GET("/content", controller.ReadContentController)        // Read (all)
	public.GET("/content/:userid/:id", controller.GetContentbyIDController) // APP - Read (Specific) using public url
	private.POST("/content/:id", controller.UpdateContentbyIDController) // working
	private.DELETE("/content/:id", controller.DeleteContentbyIDController) // Delete Content as well delete from playlist
	private.POST("/uploadcontent", controller.UploadContentController) // Upload single
	private.POST("/uploadmultiplecontent", controller.UploadMultipleContentController) // Upload multiple

}

func ContentListRouter() {
	private.POST("/contentlist", controller.CreateContentListController)
	// R.GET("/contentlist", controller.ReadContentController)
	// R.GET("/contentlist/:id", controller.GetContentbyIDController)

}

func ScreenRouter() { //Done
	myMap := make(map[string]chan DataModel.EventStreamRequest)
	// ch := make(chan DataModel.EventStreamRequest)
	private.POST("/screen", controller.CreateScreenController)
	private.GET("/screen", controller.ReadScreenController)
	private.GET("/screen/:id", controller.GetScreenbyIDController)
	private.GET("/screen/getplaylists/:id", controller.GetAllPlaylistsforScreenController)
	public.POST("/screen/getplaylists", controller.GetPlaylistofScreenController) // APP -  return the detail of a playlist of a screen
	private.POST("/screen/:id", controller.UpdateScreenbyIDController) // working
	public.POST("/screen/:id", controller.PublicUpdateScreenbyIDController) // working
	private.DELETE("/screen/:id", controller.DeleteScreenbyIDController) // Delete Screen as well delete from playlist

	private.POST("/event-stream/:id", func(c *gin.Context) {
		screencode := c.Params.ByName("id")
		_, ok := myMap[screencode]
		if !ok {
			myMap[screencode] = make(chan DataModel.EventStreamRequest)
		}
		controller.HandleEventStreamPost(c, myMap[screencode], screencode)
		delete(myMap, screencode)
	})
	public.GET("/event-stream/:id", middleware.HeadersMiddleware(), func(c *gin.Context) {
		screencode := c.Params.ByName("id")
		_, ok := myMap[screencode]
		if !ok {
			myMap[screencode] = make(chan DataModel.EventStreamRequest)
		}
		controller.HandleEventStreamGet(c, myMap[screencode])
	})
}
