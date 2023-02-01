package auth

import (
	validate "advocate-back/internal/delivery/http/validator"
	"advocate-back/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	m := new(validate.AuthRequest)
	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if m.Username != config.AppConfig.Auth.Username || m.Password != config.AppConfig.Auth.Password {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"Vladimir",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig.Auth.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
