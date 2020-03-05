package store

import (
	"errors"

	"github.com/alyyousuf7/goserver/model"
)

var (
	ErrNotFound = errors.New("not found")
)

type UserStorer interface {
	Create(*model.User) error
	FindOne(where ...interface{}) (*model.User, error)
}

type ClientStorer interface {
	Create(*model.Client) error
	FindOne(where ...interface{}) (*model.Client, error)
	AddUser(model.Client, model.User, model.ClientRole) error
}
