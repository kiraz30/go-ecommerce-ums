package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, User *models.User) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, request *models.User) (*models.User, error)
	RegisterAdmin(ctx context.Context, request *models.User) (*models.User, error)
}

type IUserHandler interface {
	RegisterUserHandler(e echo.Context) error
	RegisterAdminHandler(e echo.Context) error
}
