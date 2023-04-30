package app

import (
	"advocate-back/internal/algorithm"
	"advocate-back/internal/bot"
)

func Run() {
	//db.StartRedis()
	//httpServer := server.NewServer()
	//err := httpServer.Connect()
	//if err != nil {
	//	return
	//}
	algorithm.CheckAlgorithm(&bot.MockBot)
}
