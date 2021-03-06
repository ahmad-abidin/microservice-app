package api

import (
	"errors"
	"fmt"
	"microservice-app/auth-service/model"
	ucs "microservice-app/auth-service/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type server struct {
	usecase ucs.Usecase
}

// NewDeliveryAPI ...
func NewDeliveryAPI(u ucs.Usecase) *echo.Echo {
	s := echo.New()
	authServer := &server{
		usecase: u,
	}

	s.GET("/authentication", authServer.Authentication)
	s.GET("/authorization", authServer.Authorization)

	return s
}

// Authentication ...
func (u *server) Authentication(ctx echo.Context) error {
	base64BasicAuth := ctx.Request().Header.Get("Authorization")
	token, err := u.usecase.Authentication(base64BasicAuth)
	if err != nil {
		return ctx.JSON(http.StatusOK, fmt.Errorf("Error : %v", model.Log("e", "api-Aen_Aen", err)))
	}

	model.Log("s", "api-Aor", errors.New("Successfully Authentication"))

	return ctx.JSON(http.StatusOK, &token)
}

// Authorization ...
func (u *server) Authorization(ctx echo.Context) error {
	unsignedToken := ctx.Request().Header.Get("Authorization")
	identity, err := u.usecase.Authorization(unsignedToken)
	if err != nil {
		return ctx.JSON(http.StatusOK, fmt.Errorf("Error : %v", model.Log("e", "api-Aor_Aor", err)))
	}

	model.Log("s", "api-Aor", errors.New("Successfully Athorization"))

	return ctx.JSON(http.StatusOK, identity)
}
