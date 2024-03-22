package main

import (
	"final-project-mygram/lib"
	"final-project-mygram/router"
)

func main() {
	lib.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}