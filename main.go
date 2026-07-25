package main

import (
	"go-project/database"
	"go-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDb()

	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")

}
