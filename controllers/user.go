package controllers

import (
	"echo-sample/repository"
	"echo-sample/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Response struct {
	Result string `json:"result"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, Response{Result: "FAILED"})
	}

	repo := struct {
		repository.UserRepository
		repository.SeasonRepository
	}{}
	if err := services.CreateUser(&repo, services.CreateUserArgs{Name: "nogi", Age: 8}); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, Response{Result: "FAILED"})
	}

	return c.JSON(http.StatusCreated, Response{Result: "SUCCESS"})
}
