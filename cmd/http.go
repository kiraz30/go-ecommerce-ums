package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/repository"
	"ecommerce-ums/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {

	dependency := dependencyInject()

	healthyCheck := &api.HealthcheckAPI{}
	e := echo.New()
	e.GET("/healthcheck", healthyCheck.Healthcheck)

	userV1 := e.Group("/user/v1")
	userV1.POST("/register", dependency.UserAPI.RegisterUserHandler)
	userV1.POST("/register/admin", dependency.UserAPI.RegisterAdminHandler)
	userV1.POST("/login", dependency.UserAPI.Login)
	userV1.POST("/login/admin", dependency.UserAPI.LoginAdmin)
	userV1.PUT("/refresh-token", dependency.RefreshTokenAPI.RefreshToken, dependency.MiddlewareValidateRefreshToken)
	userV1.DELETE("/logout", dependency.UserAPI.Logout, dependency.MiddlewareValidateAuth)
	userV1.GET("/profile", dependency.UserAPI.GetProfile, dependency.MiddlewareValidateAuth)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	UserRepository  interfaces.IUserRepository
	UserAPI         interfaces.IUserHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler
}

func dependencyInject() Dependency {
	userRepository := &repository.UserRepository{
		DB: helpers.DB,
	}

	userService := &services.UserService{
		UserRepository: userRepository,
	}

	userAPI := &api.UserAPI{
		UserService: userService,
	}

	refreshTokenService := &services.RefreshTokenService{
		UserRepository: userRepository,
	}

	refreshTokenAPI := &api.RefreshTokenAPI{
		RefreshTokenService: refreshTokenService,
	}

	return Dependency{
		UserAPI:         userAPI,
		UserRepository:  userRepository,
		RefreshTokenAPI: refreshTokenAPI,
	}
}
