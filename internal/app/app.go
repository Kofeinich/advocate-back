package app

import (
	server "advocate-back/internal/delivery/http"
	"advocate-back/internal/repository/redis"
)

func Run() {
	redis.StartRedis()
	httpServer := server.NewServer()
	httpServer.Connect()
}
