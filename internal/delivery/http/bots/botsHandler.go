package bots

import (
	"advocate-back/internal/services/botService"
	"advocate-back/internal/states"
)

type BotHandler struct {
	s service
}

type service interface {
	AddBot(conf states.BotStates, token string) error
	DeleteBot(id string) error
	GetAllBots() ([]botService.Bot, error)
	UpdateBotConfig(conf states.BotStates, id string) error
}

func NewBotHandler(s *botService.Service) *BotHandler {
	return &BotHandler{s: s}
}
