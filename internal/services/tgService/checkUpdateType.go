package tgService

import (
	"bot_forge_back/internal/delivery/http/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Types string

const (
	MessageType       Types = "message"
	CallbackQueryType Types = "callback_query"
)

func CheckUpdateType(update validator.TgValidatorRequest) (Types, error) {
	if update.Message != nil {
		return MessageType, nil
	}
	if update.CallbackQuery != nil {
		return CallbackQueryType, nil
	}
	return "", echo.NewHTTPError(http.StatusBadRequest, "invalid request")
}
