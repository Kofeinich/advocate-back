package app

import (
	server "advocate-back/internal/delivery/http"
)

func Run() {
	//myMap := make(map[string]string)
	//mes := make([]map[string]string, 0, 5)
	httpServer := server.NewServer()
	httpServer.Start()
	//email.SendMessage()
}
