package services

import (
	"echo-sample/models"
)

// CreateUserで必要な関数定義
type CreateUserRepository interface {
	CreateUser(user *models.User) error
	CheckSpring() error
}

type CreateUserArgs struct {
	Name string
	Age  uint
}

func CreateUser(repo CreateUserRepository, args CreateUserArgs) *models.User {
	u := &models.User{args.Name, args.Age}
	if err := repo.CreateUser(u); err != nil {
		return nil
	}
	return u
}
