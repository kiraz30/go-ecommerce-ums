package repository

import (
	"context"
	"ecommerce-ums/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, User *models.User) error {
	return r.DB.Create(User).Error
}
