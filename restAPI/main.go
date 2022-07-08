package main

import (
	"example/go-learning/restAPI/model"
	"example/go-learning/restAPI/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.OK(service.Albums))
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	album, err := service.GetAlbumById(id)
	if err != nil {
		log.Printf("Error:%s", err.Error())
		ctx.IndentedJSON(http.StatusOK, model.NotFound())
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.OK(album))
}

func postAlbums(ctx *gin.Context) {
	var newAlbum model.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.Error(err)
		return
	}

	service.AddAlbum(newAlbum)
	ctx.JSON(http.StatusOK, model.OK(newAlbum))
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/postAlbum", postAlbums)
	router.Run("localhost:8080")
}
