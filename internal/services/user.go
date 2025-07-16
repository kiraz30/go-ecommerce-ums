package services

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, request *models.User, role string) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)
	request.Role = role

	err = s.UserRepository.InsertNewUser(ctx, request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil

}

func (s *UserService) Login(ctx context.Context, request models.LoginRequest, role string) (models.LoginReponse, error) {

	var (
		response models.LoginReponse
		now      = time.Now()
	)
	userDetail, err := s.UserRepository.GetUserByUserName(ctx, request.Username, role)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}
	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	// session
	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert new user session")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}
