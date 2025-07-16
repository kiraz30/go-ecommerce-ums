package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RefreshTokenAPI struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenAPI) RefreshToken(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	refreshToken := e.Request().Header.Get("Authorization")
	//get token form echo context
	token := e.Get("token")
	tokenClaim, ok := token.(*helpers.ClaimToken)
	if !ok {
		log.Info("Failed to parse claim tto claimToken")
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	response, err := api.RefreshTokenService.RefreshToken(e.Request().Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Info("Failed to refresh token service", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, response)

}
