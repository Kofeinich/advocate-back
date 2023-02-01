package smtp

import (
	"advocate-back/pkg/config"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Email from client</title>
	</head>
	<body>
		<p>Имя: <b>%s</b>.</p>
		<p>Email: <b>%s</b>.</p>
		<p>Телефон: <b>%s</b>.</p>
		<p>Сообщение: <b>%s</b>.</p>
	</body>
</html>`

func (s *server) SendMessage(message, emailAddress, name, phone string) error {
	email := mail.NewMSG()
	email.SetFrom(config.AppConfig.Smtp.From).
		AddTo(config.AppConfig.Smtp.To).
		SetSubject("Новое сообщение от клиента")
	formattedEmailBody := fmt.Sprintf(htmlBody, name, emailAddress, phone, message)
	email.SetBody(mail.TextHTML, formattedEmailBody)

	if email.Error != nil {
		log.Println(email.Error)
		return email.Error
	}
	err := email.Send(s.connect())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Email Sent")
	return nil
}
