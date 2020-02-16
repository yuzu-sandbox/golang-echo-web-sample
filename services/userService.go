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

func CreateUser(repo CreateUserRepository, u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := repo.CreateUser(u); err != nil {
		return err
	}
	return nil
}
