package telegram

import (
	validate "bot_forge_back/internal/delivery/http/validator"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (tg TgHandler) TgWebhook(c echo.Context) (err error) {
	var m validate.TgValidatorRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	botID := c.Param("bot_id")

	sendRequest, err := tg.s.ProcessTgUpdate(botID, m)

	bot, err := tgbotapi.NewBotAPI(sendRequest.BotToken)

	bot.Send(sendRequest.Msg)

	return c.JSON(http.StatusOK, sendRequest.Msg)
}
