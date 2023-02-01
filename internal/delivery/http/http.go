package http

import (
	"advocate-back/internal/delivery/http/auth"
	validate "advocate-back/internal/delivery/http/validator"
	"advocate-back/pkg/smtp"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	e          *echo.Echo
	smtpServer smtp.Server
}

func (s *Server) E() *echo.Echo {
	return s.e
}

func NewServer(smtpServer smtp.Server) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &validate.CustomValidator{Validator: validator.New()}
	return &Server{e: e, smtpServer: smtpServer}
}

func (s *Server) saveMessageRequest(c echo.Context) (err error) {
	m := new(validate.MailFromHTTP)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = s.smtpServer.SendMessage(m.Message, m.Email, m.Name, m.Phone)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func (s *Server) Connect() error {
	s.e.POST("/send_message", s.saveMessageRequest)
	s.e.POST("/login", auth.Login)
	s.e.Logger.Fatal(s.e.Start(":1323"))
	return nil
}
