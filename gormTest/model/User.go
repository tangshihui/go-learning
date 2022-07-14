package model

import "time"

type User struct {
	Id        int       `json:"id",gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//自定义表名
func (User) TableName() string {
	return "user"
}

type CreditCard struct {
	Id         int       `json:"id",gorm:"primaryKey"`
	UserId     string    `json:"user_id"`
	CardNumber string    `json:"card_number"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

//自定义表名
func (CreditCard) TableName() string {
	return "credit_card"
}
