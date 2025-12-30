package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/api-rest-go/internal/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/albums", handlers.HandleGetAllAlbums)
	router.GET("/albums/:id", handlers.HandleGetAlbumByID)
	router.GET("/schools", handlers.HandleGetAllSchools)
	router.POST("/schools", handlers.HandleCreateSchools)
	router.DELETE("/schools/:uuid", handlers.HandleDeleteSchool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback para local
	}

	router.Run("0.0.0.0:" + port)
}
