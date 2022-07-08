package service

import (
	"errors"
	"example/go-learning/restAPI/model"
)

var Albums = []model.Album{
	{ID: "1", Title: "Brain", Artist: "Tom", Price: 36},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 38.9},
}

/**
getAlbumById returns the album with id
*/
func GetAlbumById(id string) (model.Album, error) {
	for _, album := range Albums {
		if album.ID == id {
			return album, nil
		}
	}

	return model.Album{}, errors.New("Not found")
}

func AddAlbum(album model.Album) {
	Albums = append(Albums, album)
}
