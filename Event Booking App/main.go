package main

import (
	"event-booking-app/db"
	"event-booking-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegsterRoutes(server)

	server.Run(":8080")
}
