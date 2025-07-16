package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, User *models.User) error
	GetUserByUserName(ctx context.Context, username, role string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionToken(ctx context.Context, token string) (models.UserSession, error)
	GetUserRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
	UpdateRefreshToken(ctx context.Context, token, refreshToken string) error
	DeleteUserSession(ctx context.Context, token string) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, request *models.User, role string) (*models.User, error)
	Login(ctx context.Context, request models.LoginRequest, role string) (models.LoginReponse, error)
	GetProfile(ctx context.Context, username string) (models.User, error)
	Logout(ctx context.Context, token string) error
}

type IUserHandler interface {
	RegisterUserHandler(e echo.Context) error
	RegisterAdminHandler(e echo.Context) error
	Login(e echo.Context) error
	LoginAdmin(e echo.Context) error
	GetProfile(e echo.Context) error
	Logout(e echo.Context) error
}
