package routes

import (
	middlewares "go-project/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Auth)
	authenticated.PUT("/updateEvent/:id", updateEvent)
	authenticated.POST("/events", saveEvent)
	authenticated.DELETE("/delete/:id", deleteEvent)

	server.POST("/signup", createUser)
	server.POST("/login", loginUser)

}
