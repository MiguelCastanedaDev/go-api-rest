package services

import (
	"github.com/api-rest-go/internal/types"
)

func GetAllAlbumsService() []types.Album {

	var albums = []types.Album{
		{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
		{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1957},
		{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
	}

	return albums
}
