package bots

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h BotHandler) GetAllBots(c echo.Context) (err error) {
	bots, err := h.s.GetAllBots()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, bots)
}
