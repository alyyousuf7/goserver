package main

import (
	"fmt"

	"github.com/alyyousuf7/goserver/app"
	"github.com/alyyousuf7/goserver/model"
	"github.com/alyyousuf7/goserver/store"
)

func main() {
	db, err := model.Setup("localhost", "postgres", "postgres", "postgres")
	if err != nil {
		panic(err)
	}

	userStore := store.NewUserStore(db)
	clientStore := store.NewClientStore(db)
	app := &app.App{
		UserStore:   userStore,
		ClientStore: clientStore,
		TokenSecret: []byte("HelloWorld"),
	}

	// Seed data
	users := []*model.User{
		{
			Email:    "admin1",
			Password: "HelloWorld",
			Role:     model.RoleAdmin,
		},
	}
	for _, u := range users {
		if err := userStore.Create(u); err != nil {
			panic(err)
		}
	}

	// Application
	token, err := app.GenerateToken("admin1", "HelloWorld")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
