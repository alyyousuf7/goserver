package app

import (
	"errors"

	"github.com/alyyousuf7/goserver/model"
	"github.com/alyyousuf7/goserver/store"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates token for the given email and password
func (app *App) GenerateToken(email, password string) (string, error) {
	user, err := app.UserStore.FindOne(model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return "", ErrAuthFailed
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"role":   user.Role,
	})

	return token.SignedString(app.TokenSecret)
}

// VerifyToken checks the token validity
func (app *App) VerifyToken(tokenStr string) error {
	_, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrAuthFailed
		}

		return app.TokenSecret, nil
	})
	if err != nil {
		return err
	}

	return nil
}
