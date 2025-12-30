package main

import (
	"github.com/gin-gonic/gin"

	"github.com/api-rest-go/internal/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/albums", handlers.HandleGetAllAlbums)

	router.Run("localhost:8080")
}
