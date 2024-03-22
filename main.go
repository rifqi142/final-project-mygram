package main

import (
	"final-project-mygram/config"
	"final-project-mygram/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}