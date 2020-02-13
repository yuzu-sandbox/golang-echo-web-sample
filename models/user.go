package models

import "errors"

type User struct {
	Name string
	Age  uint
}

func (u *User) WrappedName() string {
	return "[" + u.Name + "]"
}

func (u *User) Validate() error {
	if u.Age <= 18 {
		return errors.New("User is not minor")
	}
	return nil
}
