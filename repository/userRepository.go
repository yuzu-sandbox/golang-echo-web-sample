package repository

import (
	"echo-sample/models"
	"errors"
)

var db = make(map[string]models.User)

//var _ services.CreateUserRepository = (*UserRepository)(nil)

type UserRepository struct{}

func (_ *UserRepository) CreateUser(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if _, ok := db[u.Name]; ok {
		return errors.New("User is already exists")
	}

	db[u.Name] = *u
	return nil
}
