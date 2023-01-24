package validator

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	MailFromHTTP struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Phone   string `json:"phone" validate:"required,phone"`
		Message string `json:"message" validate:"required"`
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
