package cmd

import (
	"my-echo-framework/helpers"
	"my-echo-framework/internal/api"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {

	healthyCheck := &api.HealthcheckAPI{}
	e := echo.New()
	e.GET("/healthcheck", healthyCheck.Healthcheck)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}
