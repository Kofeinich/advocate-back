package validator

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	SendMessageRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Phone   string `json:"phone" validate:"required"`
		Message string `json:"message" validate:"required"`
	}

	AuthRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
