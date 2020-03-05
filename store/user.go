package store

import (
	"github.com/alyyousuf7/goserver/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

var _ UserStorer = &UserStore{}

func NewUserStore(db *gorm.DB) UserStorer {
	return &UserStore{db}
}

func (u *UserStore) Create(user *model.User) error {
	query := u.db.Model(model.User{}).Create(user)
	if err := queryError(query); err != nil {
		return err
	}

	return nil
}

func (u *UserStore) FindOne(where ...interface{}) (*model.User, error) {
	user := model.User{}

	query := u.db.Model(model.User{}).First(&user, where...)
	if err := queryError(query); err != nil {
		return nil, err
	}

	return &user, nil
}
