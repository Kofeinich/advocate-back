package telegram

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (tg TgHandler) TgWebhook(c echo.Context) (err error) {
	//botId := c.Param("bot_id")

	// todo bind context into struct update
	// todo call update

	return c.JSON(http.StatusOK, "ok")

}
