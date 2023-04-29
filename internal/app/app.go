package app

import (
	"advocate-back/internal/algorithm"
	"advocate-back/internal/bot"
	server "advocate-back/internal/delivery/http"
	"advocate-back/pkg/db"
)

func Run() {
	db.StartRedis()
	httpServer := server.NewServer()
	httpServer.Connect()
	algorithm.CheckAlgorithm(bot.MockBot.BotStates)
}
