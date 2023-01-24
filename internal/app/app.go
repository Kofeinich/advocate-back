package app

import (
	server "advocate-back/internal/delivery/http"
)

func Run() {
	httpServer := server.NewServer()
	httpServer.Start()
	//email.SendMessage()
}
