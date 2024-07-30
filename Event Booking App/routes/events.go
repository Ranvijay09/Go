package routes

import (
	"event-booking-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}
	event, err := models.GetEventByID(int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	allEvents, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events!"})
		return
	}
	context.JSON(http.StatusOK, allEvents)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something Went Wrong!"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}
