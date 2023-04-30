package bots

import (
	validate "advocate-back/internal/delivery/http/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h BotHandler) AddBot(c echo.Context) (err error) {
	m := new(validate.AddBotRequest)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return err
	}
	//algorithm.CheckAlgorithm(&bot.MockBot)
	h.s.AddBot(m.BotConfig, m.Token)
	return c.JSON(http.StatusOK, m)
}
