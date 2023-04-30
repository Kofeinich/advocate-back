package app

import (
	"advocate-back/internal/delivery/http"
	"advocate-back/internal/delivery/http/bots"
	"advocate-back/internal/repository/botRepository"
	"advocate-back/internal/services/botService"
	"advocate-back/pkg/rdb"
)

func Run() {
	redis := rdb.StartRedis()
	repo := botRepository.NewRepo(redis)
	service := botService.NewService(repo)
	handler := bots.NewBotHandler(service)
	httpServer := http.NewServer(handler)
	err := httpServer.Connect()
	if err != nil {
		return
	}
}
