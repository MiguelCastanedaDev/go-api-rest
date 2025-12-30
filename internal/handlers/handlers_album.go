package handlers

import (
	"net/http"

	"github.com/api-rest-go/internal/services"
	"github.com/gin-gonic/gin"
)

func HandleGetAllAlbums(c *gin.Context) {
	albums := services.GetAllAlbumsService()

	if len(albums) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Whoops! No albums found.", "status": http.StatusNotFound})
		return
	}

	if albums == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops! Something went wrong while fetching albums.", "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums, "status": http.StatusOK})
}

func HandleGetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album := services.GetAlbumByIDService(id)

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album not found", "status": http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album, "status": http.StatusOK})
}
