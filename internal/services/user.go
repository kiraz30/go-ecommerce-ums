package services

import (
	"context"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, request *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)

	request.Role = "customer"

	err = s.UserRepository.InsertNewUser(ctx, request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil

}

func (s *UserService) RegisterAdmin(ctx context.Context, request *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)

	request.Role = "admin"

	err = s.UserRepository.InsertNewUser(ctx, request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil

}
