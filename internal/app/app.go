package app

import (
	"bot_forge_back/internal/delivery/http"
	"bot_forge_back/internal/delivery/http/bots"
	"bot_forge_back/internal/delivery/http/telegram"
	"bot_forge_back/internal/repository/botRepository"
	"bot_forge_back/internal/repository/tgRepository"
	"bot_forge_back/internal/services/botService"
	"bot_forge_back/internal/services/tgService"
	"bot_forge_back/pkg/rdb"
)

func Run() {
	redis := rdb.StartRedis()
	repo := botRepository.NewRepo(redis)
	userRepo := tgRepository.NewUserRepo(redis)
	service := botService.NewService(repo)
	serviceTg := tgService.NewService(userRepo, repo)
	handler := bots.NewBotHandler(service)
	handlerTg := telegram.NewTgHandler(serviceTg)
	httpServer := http.NewServer(handler, handlerTg)
	list, err := repo.GelAllBotsFromList()
	if err != nil {
		return
	}
	for _, botID := range list {
		botToken, err := repo.GetBotTokenByID(botID)
		if err != nil {
			continue
		}
		err = tgService.RegNewWebHook(botToken, botID)
		if err != nil {
			continue
		}
	}

	err = httpServer.Connect()
	if err != nil {
		return
	}
}
