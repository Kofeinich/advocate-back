package tgService

import (
	"advocate-back/internal/delivery/http/validator"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type SendMessageRequest struct {
	Update   validator.TgValidatorRequest
	BotID    string
	BotToken string
	Msg      tgbotapi.MessageConfig
}
