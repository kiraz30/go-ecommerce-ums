package cmd

import (
	"ecommerce-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		token := e.Request().Header.Get("Authorization")
		if token == "" {
			log.Println("Authorization header is missing")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized", nil)

		}

		if d.UserRepository == nil {
			log.Println("UserRepository is not initialized")
			return helpers.SendResponseHTTP(e, http.StatusInternalServerError, "Internal Server Error", nil)

		}

		_, err := d.UserRepository.GetUserSessionToken(e.Request().Context(), token)
		if err != nil {
			log.Println("Error GET user session token:", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized validate token", nil)
		}

		claim, err := helpers.ValidateToken(e.Request().Context(), token)
		if err != nil {
			log.Println("Error validating token:", err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized validate token", nil)

		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("Token has expired", claim.ExpiresAt)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized token expired", nil)

		}
		e.Set("token", claim)
		return next(e)
	}
}

func (d *Dependency) MiddlewareValidateRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		auth := e.Request().Header.Get("Authorization")
		if auth == "" {
			log.Println("Authorization header is empty")
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized", nil)
		}

		_, err := d.UserRepository.GetUserRefreshToken(e.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized", nil)

		}

		claim, err := helpers.ValidateToken(e.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Unauthorized validate token", nil)

		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("Token expired", claim.ExpiresAt)
			return helpers.SendResponseHTTP(e, http.StatusUnauthorized, "Token expired", nil)

		}
		e.Set("token", claim)
		return next(e)
	}

}
