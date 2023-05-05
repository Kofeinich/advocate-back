package telegram

import (
	validate "advocate-back/internal/delivery/http/validator"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (tg TgHandler) TgWebhook(c echo.Context) (err error) {
	m := &validate.TgValidatorRequest{}

	if err = json.NewDecoder(c.Request().Body).Decode(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	botID := c.Param("bot_id")

	_, err = tg.s.ProcessTgUpdate(botID, m.Update)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}
