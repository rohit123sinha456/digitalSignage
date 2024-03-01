package main

import (
	"github.com/rohit123sinha456/digitalSignage/dbmaster"
)

func main() {
	client := dbmaster.ConnectDB()
	dbmaster.CreateUser(client, "Rohit Sinha")
}
