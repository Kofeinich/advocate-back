package telegram

import (
	"advocate-back/internal/services/tgService"
)

type TgHandler struct {
	s service
}

type service interface {
}

func NewTgHandler(s *tgService.Service) *TgHandler {
	return &TgHandler{s: s}
}
