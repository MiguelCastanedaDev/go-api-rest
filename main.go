package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1957},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
}

func handleGetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func handlePostAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func handleGetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", handleGetAlbums)
	router.POST("/albums", handlePostAlbums)
	router.GET("/albums/:id", handleGetAlbumById)

	router.Run("localhost:8080")
}
