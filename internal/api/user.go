package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	UserService interfaces.IUserService
}

func (api *UserAPI) RegisterUserHandler(e echo.Context) error {

	var (
		log = helpers.Logger
	)

	request := models.User{}

	err := e.Bind(&request)
	if err != nil {
		log.Info("Error binding request :", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err = request.Validate(); err != nil {
		log.Info("Error validate user register request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	response, err := api.UserService.RegisterUser(e.Request().Context(), &request)
	if err != nil {
		log.Info("Failed to register user :", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, response)
}

func (api *UserAPI) RegisterAdminHandler(e echo.Context) error {

	var (
		log = helpers.Logger
	)

	request := models.User{}

	err := e.Bind(&request)
	if err != nil {
		log.Info("Error binding request :", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err = request.Validate(); err != nil {
		log.Info("Error validate user register request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	response, err := api.UserService.RegisterAdmin(e.Request().Context(), &request)
	if err != nil {
		log.Info("Failed to register user :", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, response)
}
