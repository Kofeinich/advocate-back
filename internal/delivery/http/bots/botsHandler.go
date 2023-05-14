package bots

import (
	"bot_forge_back/internal/services/botService"
	"bot_forge_back/internal/states"
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
