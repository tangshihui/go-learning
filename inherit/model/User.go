package model

import (
	"fmt"
	"time"
)

type ModelBase struct {
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (m ModelBase) String() string {
	return fmt.Sprintf("%v,%v", m.CreatedTime.Format(time.RFC3339), m.UpdatedTime.Format(time.RFC3339))
}

type User struct {
	ModelBase
	Id      int
	Name    string
	Address string
}

func (u User) String() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", u.CreatedTime.Format(time.RFC3339), u.UpdatedTime.Format(time.RFC3339), u.Id, u.Name, u.Address)
}
