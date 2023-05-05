package telegram

import (
	"advocate-back/internal/services/tgService"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgHandler struct {
	s service
}

type service interface {
	ProcessTgUpdate(botId string, update tgbotapi.Update) (blockId string, err error)
}

func NewTgHandler(s *tgService.Service) *TgHandler {
	return &TgHandler{s: s}
}
