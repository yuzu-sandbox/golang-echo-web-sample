package repository

import "errors"

type SeasonRepository struct {
}

func (_ *SeasonRepository) CheckSpring() error {
	return errors.New("Not implemented")
}
