package main

import (
	"example/go-learning/restAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums = []model.Album{
	{ID: "1", Title: "Brain", Artist: "Tom", Price: 36},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 38.9},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.OK(albums))
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, model.OK(album))
			return
		}
	}

	ctx.IndentedJSON(http.StatusOK, model.NotFound())
}

func postAlbums(ctx *gin.Context) {
	var newAlbum model.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.Error(err)
		return
	}

	albums = append(albums, newAlbum)
	ctx.JSON(http.StatusOK, model.OK(newAlbum))
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/postAlbum", postAlbums)
	router.Run("localhost:8080")
}
