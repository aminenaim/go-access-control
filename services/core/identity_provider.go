package core

import (
	"core/models"
	"fmt"
)

type IdentityProvider struct {
	Users []models.User
}

type IdentityProviderInterface interface {
	GetUserFromID(id string) (models.User, error)
}

func NewMockIdentityProvider() IdentityProvider {
	return IdentityProvider{
		Users: []models.User{
			{
				ID:   "user1",
				Name: "John",
				Role: "user",
			},
			{
				ID:   "user2",
				Name: "Jane",
				Role: "admin",
			},
		},
	}
}

func (i IdentityProvider) GetUserFromID(id string) (models.User, error) {
	for _, user := range i.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}
