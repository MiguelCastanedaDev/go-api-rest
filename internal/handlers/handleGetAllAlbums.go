package handlers

import (
	"net/http"

	"github.com/api-rest-go/internal/services"
	"github.com/gin-gonic/gin"
)

func HandleGetAllAlbums(c *gin.Context) {
	albums := services.GetAllAlbumsService()
	c.JSON(http.StatusOK, albums)
}
