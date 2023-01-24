package app

import (
	server "advocate-back/internal/delivery/http"
	"advocate-back/internal/repository/redis"
)

func Run() {
	httpServer := server.NewServer()
	httpServer.Start()
	redis.StartRedis()
	//email.SendMessage()
}
