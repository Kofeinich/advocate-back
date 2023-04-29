package http

import (
	"advocate-back/internal/delivery/http/auth"
	validate "advocate-back/internal/delivery/http/validator"
	config2 "advocate-back/pkg"
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	e *echo.Echo
}

func (s *Server) E() *echo.Echo {
	return s.e
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &validate.CustomValidator{Validator: validator.New()}
	return &Server{e: e}
}

func (s *Server) saveMessageRequest(c echo.Context) (err error) {
	m := new(validate.SendMessageRequest)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func (s *Server) Connect() error {
	s.e.Use(middleware.CORS())
	s.e.POST("/send_message", s.saveMessageRequest)
	s.e.POST("/login", auth.Login)
	s.e.POST("/refresh", auth.Refresh)
	g := s.e.Group("/restricted")
	config := echojwt.Config{
		SigningKey: []byte(config2.AppConfig.Auth.Secret),
	}
	g.Use(echojwt.WithConfig(config))
	s.e.Logger.Fatal(s.e.Start(":1323"))
	return nil
}
