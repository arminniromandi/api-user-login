package routes

import (
	"go-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	data, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cloud not get data TryAgain!"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func getEvent(c *gin.Context) {
	//10= Decimal
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not get event id!"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "there was error in saving data in database!"})
		return
	}

	c.JSON(http.StatusFound, event)

}

func saveEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error in data " + err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")
	event.User_Id = userId
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "there was an error On saving Data! Try Again!",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "saved",
		"data":   event,
	})

}

func updateEvent(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not get event id!"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not fetch event!"})
		return
	}

	var updateEvent models.Event

	err = c.ShouldBindJSON(&updateEvent)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error in data " + err.Error(),
		})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not find data to update!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data update Successfuly "})

}

func deleteEvent(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not get event id!"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not fetch event!"})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cloud not delete event!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event update Successfuly "})

}
