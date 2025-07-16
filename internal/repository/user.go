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

func (r *UserRepository) GetUserRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("refresh_token =?", refreshToken).Last(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("user session not found")
	}
	return session, nil
}

func (r *UserRepository) UpdateRefreshToken(ctx context.Context, token, refreshToken string) error {
	return r.DB.Exec("Update user_session SET token = ? where refresh_token = ?", token, refreshToken).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_session WHERE token = ?", token).Error
}
