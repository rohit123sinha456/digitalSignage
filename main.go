package main

import (
	"github.com/rohit123sinha456/digitalSignage/router"
)

func main() {
	router.SetupRouter()
	router.UserRouter()
	router.DeviceRouter()
	router.PlaylistRouter()
	router.ContentRouter()
	router.ContentListRouter()
	router.R.Run(":8080")
}
