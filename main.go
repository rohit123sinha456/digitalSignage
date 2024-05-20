package main

import (
	"github.com/rohit123sinha456/digitalSignage/router"
)

func main() {
	router.SetupRouter()
	router.UserRouter()
	router.PlaylistRouter()
	router.ContentRouter()
	router.ContentListRouter()
	router.ScreenRouter()
	router.AuthRoutes()
	router.R.Run(":8080")
}
