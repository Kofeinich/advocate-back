package bots

import (
	validate "bot_forge_back/internal/delivery/http/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h BotHandler) DeleteBot(c echo.Context) (err error) {
	m := new(validate.DeleteBotRequest)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return err
	}
	err = h.s.DeleteBot(m.BotID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}
