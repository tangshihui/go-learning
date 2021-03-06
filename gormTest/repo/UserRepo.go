package repo

import (
	"example/go-learning/gormTest/model"
	"gorm.io/gorm"
	"log"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// list all users by paganation
//
//page
//
//perPage
func (r *UserRepo) ListAllUsers(page, perPage int) ([]model.User, error) {
	users := make([]model.User, 0)
	offset := (page - 1) * perPage
	result := r.db.Model(&model.User{}).Order("name desc").Limit(perPage).Offset(offset).Find(&users)

	if result.Error != nil {
		log.Printf("failed set user")
	}

	return users, nil
}

func (r *UserRepo) SaveUser(user model.User) (*model.User, error) {
	res := r.db.Create(&user)

	log.Printf("inserted user:%v", user.Id)

	return &user, res.Error
}

func (r *UserRepo) DeleteUser(id int) (int64, error) {
	res := r.db.Delete(&model.User{}, id)
	return res.RowsAffected, res.Error
}

func (r *UserRepo) UpdateUser(user *model.User) (int64, error) {
	res := r.db.Model(&model.User{}).Where("id=?", user.Id).Update("name", user.Name).Update("address", user.Address)

	return res.RowsAffected, res.Error
}

func (r *UserRepo) QueryUser(id int) (*model.User, error) {
	user := model.User{}
	find := r.db.First(&user, id)

	if find.Error != nil {
		log.Printf("failed find user by id:%v", id)
		return nil, find.Error
	}

	return &user, nil
}
