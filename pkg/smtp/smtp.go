package smtp

import (
	"advocate-back/pkg/config"
	"crypto/tls"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"time"
)

type Server interface {
	SendMessage(string, string, string, string) error
}

type server struct {
	s *mail.SMTPServer
}

func NewServer() *server {
	c := mail.NewSMTPClient()
	c.Host = config.AppConfig.Smtp.Host
	c.Port = config.AppConfig.Smtp.Port
	c.Username = config.AppConfig.Smtp.Username
	c.Password = config.AppConfig.Smtp.Password
	c.Encryption = mail.EncryptionSTARTTLS
	c.KeepAlive = false
	c.ConnectTimeout = 10 * time.Second
	c.SendTimeout = 10 * time.Second
	c.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &server{s: c}
}

func (s *server) connect() *mail.SMTPClient {
	smtpClient, err := s.s.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return smtpClient
}
