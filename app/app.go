package app

import "github.com/alyyousuf7/goserver/store"

type App struct {
	UserStore   store.UserStorer
	ClientStore store.ClientStorer

	TokenSecret []byte // TODO: I don't know, may be use JWT interface?
}
