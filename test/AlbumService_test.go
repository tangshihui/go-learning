package test

import (
	"example/go-learning/restAPI/model"
	"example/go-learning/restAPI/service"
	"testing"
)

func TestGetAlbumById(t *testing.T) {

	var id = "1"
	actual, _ := service.GetAlbumById(id)
	expected := model.Album{ID: "1", Title: "Brain", Artist: "Tom", Price: 36}

	if actual != expected {
		t.Errorf("actual:%v not equals to expected:%v", actual, expected)
	}
}
