package repository

import (
	"context"
	"ecommerce-ums/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, User *models.User) error {
	return r.DB.Create(User).Error
}

func (r *UserRepository) GetUserByUserName(ctx context.Context, username, role string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	sql := r.DB.Where("username = ?", username)

	if role != "" {
		sql = sql.Where("role = ?", role)
	}

	err = sql.First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) GetUserSessionToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		dataUserSession models.UserSession
		err             error
	)

	err = r.DB.Where("token = ?", token).First(&dataUserSession).Error
	if err != nil {
		return dataUserSession, err
	}
	if dataUserSession.ID == 0 {
		return dataUserSession, errors.New("user session not found")
	}

	return dataUserSession, nil
}
