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
	response, err := api.UserService.RegisterUser(e.Request().Context(), &request, constants.RoleCustomer)
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
	response, err := api.UserService.RegisterUser(e.Request().Context(), &request, constants.RoleAdmin)
	if err != nil {
		log.Info("Failed to register user :", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, response)
}

func (api *UserAPI) Login(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	request := models.LoginRequest{}
	resp := models.LoginReponse{}

	if err := e.Bind(&request); err != nil {
		log.Info("Failded to parse requst:", err)
		helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)

	}
	if err := request.Validate(); err != nil {
		log.Info("Failded to validate:", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)

	}

	resp, err := api.UserService.Login(e.Request().Context(), request, constants.RoleCustomer)
	if err != nil {
		log.Info("Failded on login service:", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)

}

func (api *UserAPI) LoginAdmin(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	request := models.LoginRequest{}
	resp := models.LoginReponse{}

	if err := e.Bind(&request); err != nil {
		log.Info("Failded to parse requst:", err)
		helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)

	}
	if err := request.Validate(); err != nil {
		log.Info("Failded to validate:", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)

	}

	resp, err := api.UserService.Login(e.Request().Context(), request, constants.RoleAdmin)
	if err != nil {
		log.Info("Failded on login service:", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) GetProfile(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Get("token")
	tokenClaim, ok := token.(*helpers.ClaimToken)
	if !ok {
		log.Info("Failded get token data")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	resp, err := api.UserService.GetProfile(e.Request().Context(), tokenClaim.Username)
	if err != nil {
		log.Info("Failded to get Profile:", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) Logout(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	token := e.Request().Header.Get("Authorization")
	err := api.UserService.Logout(e.Request().Context(), token)
	if err != nil {
		log.Info("Failded on logout service:", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, nil)
}
