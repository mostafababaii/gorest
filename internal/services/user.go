package services

import (
	"errors"
	"github.com/mostafababaii/gorest/internal/cache"
	"github.com/mostafababaii/gorest/internal/models"
)

var UserService = userService{}

type userService struct{}

func (us *userService) FindByID(userID uint) (*models.User, error) {
	user, _ := cache.UserCache.FindByID(userID)
	if user != nil {
		return user, nil
	}

	user, err := models.UserModel.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	cache.UserCache.Persist(*user)

	return user, nil
}
