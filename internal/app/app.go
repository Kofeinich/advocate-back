package app

import (
	server "advocate-back/internal/delivery/http"
	"advocate-back/internal/repository/posts"
	"advocate-back/internal/repository/redis"
	"advocate-back/pkg/db"
	"advocate-back/pkg/smtp"
	"log"
)

func Run() {
	redis.StartRedis()
	postgres, err := db.NewPostgres()
	if err != nil {
		log.Fatal("Postgres dont working", err)
	}
	postsRepository, err := posts.NewPostsRepository(postgres)
	if err != nil {
		log.Fatal("repository not initialized", err)
	}

	smtpServer := smtp.NewServer()
	httpServer := server.NewServer(smtpServer)
	httpServer.Connect()
}
