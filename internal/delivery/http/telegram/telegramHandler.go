package telegram

import (
	"advocate-back/internal/delivery/http/validator"
	"advocate-back/internal/services/tgService"
)

type TgHandler struct {
	s service
}

type service interface {
	ProcessTgUpdate(botId string, update validator.TgValidatorRequest) (sendRequest tgService.SendMessageRequest, err error)
}

func NewTgHandler(s *tgService.Service) *TgHandler {
	return &TgHandler{s: s}
}
