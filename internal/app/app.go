package app

import (
	"advocate-back/internal/delivery/http"
	"advocate-back/internal/delivery/http/bots"
	"advocate-back/internal/delivery/http/telegram"
	"advocate-back/internal/repository/botRepository"
	"advocate-back/internal/repository/tgRepository"
	"advocate-back/internal/services/botService"
	"advocate-back/internal/services/tgService"
	"advocate-back/pkg/rdb"
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
