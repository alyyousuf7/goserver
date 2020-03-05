package app_test

import (
	"errors"
	"testing"

	"github.com/alyyousuf7/goserver/app"
	"github.com/alyyousuf7/goserver/mocks"
	"github.com/alyyousuf7/goserver/model"
	"github.com/alyyousuf7/goserver/store"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateToken(t *testing.T) {
	userStore := new(mocks.UserStorer)

	testApp := app.App{
		UserStore:   userStore,
		TokenSecret: []byte("Hi there"),
	}

	// Should generate token
	user := model.User{
		Model: gorm.Model{ID: 1},
		Role:  model.RoleAdmin,
	}
	userStore.ExpectedCalls = []*mock.Call{}
	userStore.On("FindOne", mock.Anything).Return(&user, nil)
	token1, err := testApp.GenerateToken("user", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token1)

	// Should generate a different token (change role)
	user = model.User{
		Model: gorm.Model{ID: 1},
		Role:  model.RoleClient,
	}
	userStore.ExpectedCalls = []*mock.Call{}
	userStore.On("FindOne", mock.Anything).Return(&user, nil)
	token2, err := testApp.GenerateToken("user", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token2)
	assert.NotEqual(t, token1, token2)

	// Should generate a different token (change id)
	user = model.User{
		Model: gorm.Model{ID: 2},
		Role:  model.RoleAdmin,
	}
	userStore.ExpectedCalls = []*mock.Call{}
	userStore.On("FindOne", mock.Anything).Return(&user, nil)
	token3, err := testApp.GenerateToken("user", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token3)
	assert.NotEqual(t, token1, token3)
	assert.NotEqual(t, token2, token3)

	// Should return auth error if user not found
	userStore.ExpectedCalls = []*mock.Call{}
	userStore.On("FindOne", mock.Anything).Return(nil, store.ErrNotFound)
	_, err = testApp.GenerateToken("user", "password")
	assert.Equal(t, app.ErrAuthFailed, err)

	// Should not return auth error if some other error occurs
	userStore.ExpectedCalls = []*mock.Call{}
	userStore.On("FindOne", mock.Anything).Return(nil, errors.New("some db error"))
	_, err = testApp.GenerateToken("user", "password")
	assert.Error(t, err)
	assert.NotEqual(t, app.ErrAuthFailed, err)
}
