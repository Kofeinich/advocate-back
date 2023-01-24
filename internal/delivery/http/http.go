package http

import (
	validate "advocate-back/internal/delivery/http/validator"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type server struct {
	e *echo.Echo
}

func (s *server) E() *echo.Echo {
	return s.e
}

func NewServer() *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &validate.CustomValidator{Validator: validator.New()}
	return &server{e: e}
}
func (s *server) Start() error {
	s.e.POST("/send_message", func(c echo.Context) (err error) {
		m := new(validate.MailFromHTTP)
		if err = c.Bind(m); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(m); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, m)
	})
	s.e.Logger.Fatal(s.e.Start(":8080"))
	return nil
}
