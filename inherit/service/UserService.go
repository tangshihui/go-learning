package service

import (
	"example/go-learning/inherit/model"
	"time"
)

func GetAllUsers() []model.User {
	users := make([]model.User, 0)
	users = append(users, model.User{
		ModelBase: model.ModelBase{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		Id:      1,
		Name:    "Tom",
		Address: "BJ",
	})

	now := time.Now()

	users = append(users, model.User{
		ModelBase: model.ModelBase{
			CreatedTime: now.Add(time.Minute * 30),
			UpdatedTime: now.Add(time.Minute * 60),
		},
		Id:      2,
		Name:    "JACK",
		Address: "SH",
	})

	return users
}
