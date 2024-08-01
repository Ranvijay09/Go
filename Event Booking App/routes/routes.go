package routes

import (
	"event-booking-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegsterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:id", updateEvent)
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("/events/:id/register", registerForEvent)
	authenticatedGroup.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
