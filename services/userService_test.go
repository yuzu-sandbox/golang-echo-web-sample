package services

import (
	"echo-sample/mock_services"
	"echo-sample/models"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := &models.User{Name:"test", Age: 24}

	mockRepo := mock_services.NewMockCreateUserRepository(ctrl)
	mockRepo.EXPECT().CreateUser(u).Return(nil)

	if err := CreateUser(mockRepo, u); err != nil {
		t.Errorf("User Create Error")
	}
}

func TestCreateUser_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := &models.User{Name:"test", Age: 14}
	mockRepo := mock_services.NewMockCreateUserRepository(ctrl)
	if err := CreateUser(mockRepo, u); err == nil {
		t.Errorf("Not got error")
	}
}

