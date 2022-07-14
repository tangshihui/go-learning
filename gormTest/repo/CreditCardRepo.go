package repo

import (
	"example/go-learning/gormTest/model"
	"gorm.io/gorm"
	"log"
)

type CreditCardRepo struct {
	db *gorm.DB
}

func NewCreditCardRepo(db *gorm.DB) *CreditCardRepo {
	return &CreditCardRepo{
		db: db,
	}
}

// list all users by paganation
//
//page
//
//perPage
func (r *CreditCardRepo) ListAll(page, perPage int) ([]model.CreditCard, error) {
	cards := make([]model.CreditCard, 0)
	offset := (page - 1) * perPage
	result := r.db.Model(&model.CreditCard{}).Limit(perPage).Offset(offset).Find(&cards)

	if result.Error != nil {
		log.Printf("failed query credit cards")
	}

	return cards, nil
}

func (r *CreditCardRepo) SaveCreditCard(card *model.CreditCard) (*model.CreditCard, error) {
	res := r.db.Create(card)

	log.Printf("inserted user:%v", card.Id)

	return card, res.Error
}

//func (r *UserRepo) DeleteUser(id int) (int64, error) {
//	res := r.db.Delete(&model.User{}, id)
//	return res.RowsAffected, res.Error
//}

//func (r *UserRepo) UpdateUser(user *model.User) (int64, error) {
//	res := r.db.Model(&model.User{}).Where("id=?", user.Id).Update("name", user.Name).Update("address", user.Address)
//
//	return res.RowsAffected, res.Error
//}

func (r *CreditCardRepo) QueryById(id int) (*model.CreditCard, error) {
	card := model.CreditCard{}
	find := r.db.First(&card, id)

	if find.Error != nil {
		log.Printf("failed find credit card by id:%v", id)
		return nil, find.Error
	}

	return &card, nil
}

func (r *CreditCardRepo) QueryByUserId(uid int) ([]model.CreditCard, error) {
	cards := make([]model.CreditCard, 0)
	find := r.db.Model(&model.CreditCard{}).Where("userId = ?", uid).Scan(&cards)

	if find.Error != nil {
		log.Printf("failed find credit card by user id:%v", uid)
		return nil, find.Error
	}

	return cards, nil
}
