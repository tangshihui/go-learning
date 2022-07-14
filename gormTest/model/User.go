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
	Id         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"user_id" gorm:"column:userId"`
	CardNumber string    `json:"card_number" gorm:"column:cardNumber"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updatedAt"`
}

//自定义表名
func (CreditCard) TableName() string {
	return "credit_card"
}
