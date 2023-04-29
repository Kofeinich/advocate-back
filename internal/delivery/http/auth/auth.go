package auth

import (
	validate "advocate-back/internal/delivery/http/validator"
	"advocate-back/pkg"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	jwt.RegisteredClaims
}

type refreshJwtCustomClaims struct {
	AccessToken string `json:"access_token"`
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

	if m.Username != pkg.AppConfig.Auth.Username || m.Password != pkg.AppConfig.Auth.Password {
		return echo.ErrUnauthorized
	}
	return generateToken(c)
}

func generateToken(c echo.Context) error {

	claims := &jwtCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(pkg.AppConfig.Auth.Secret))
	if err != nil {
		return err
	}
	refreshClaims := &refreshJwtCustomClaims{
		t,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	// Create token with claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Generate encoded token and send it as response.
	rt, err := refreshToken.SignedString([]byte(pkg.AppConfig.Auth.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  t,
		"refresh_token": rt,
	})
}

func Refresh(c echo.Context) error {
	m := new(validate.RefreshRequest)
	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var claims refreshJwtCustomClaims
	token, err := jwt.ParseWithClaims(m.RefreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(pkg.AppConfig.Auth.Secret), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if claims.AccessToken != m.AccessToken {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return generateToken(c)
}
