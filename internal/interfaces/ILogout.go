package interfaces

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}

type ILogoutHandler interface {
	Logout(e echo.Context) error
}
