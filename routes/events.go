package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch events, try again later"})
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse  event id"})
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't find event, try again later"})
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messsage": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create event, try again later"})

	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse  event id"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch  the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messsage": "Could not parse request data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse  event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messsage": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
