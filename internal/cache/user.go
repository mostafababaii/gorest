package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/mostafababaii/gorest/internal/database/redis"
	"github.com/mostafababaii/gorest/internal/models"
)

var UserCache = userCache{}

type userCache struct{}

func (uc *userCache) FindByID(userID uint) (*models.User, error) {
	userData, _ := redis.DefaultClient.HGetAll(context.Background(), fmt.Sprintf("USER:%d", userID)).Result()
	if _, ok := userData["fist_name"]; !ok {
		return nil, errors.New("user not found")
	}

	user := &models.User{
		FistName: userData["fist_name"],
		LastName: userData["last_name"],
		Email:    userData["email"],
		Password: userData["password"],
	}

	return user, nil
}

func (uc *userCache) Persist(user models.User) {
	redis.DefaultClient.HSet(context.Background(), fmt.Sprintf("USER:%d", user.ID), user)
}
