package app

import (
	server "advocate-back/internal/delivery/http"
	"advocate-back/internal/repository/redis"
	"advocate-back/pkg/smtp"
)

func Run() {
	redis.StartRedis()
	smtpServer := smtp.NewServer()
	httpServer := server.NewServer(smtpServer)
	httpServer.Connect()
}
