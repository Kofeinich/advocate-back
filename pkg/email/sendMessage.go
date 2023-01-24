package email

import (
	"crypto/tls"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"time"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

func SendMessage() {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "smtp.jino.ru"
	server.Port = 587
	server.Username = "vladimir@advocate-etalon.ru"
	server.Password = "vostok12"
	server.Encryption = mail.EncryptionSTARTTLS

	server.KeepAlive = false

	server.ConnectTimeout = 10 * time.Second

	server.SendTimeout = 10 * time.Second

	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom("vladimir@advocate-etalon.ru").
		AddTo("sultangirova_d@mail.ru").
		SetSubject("New Go Email")
	email.SetBody(mail.TextHTML, htmlBody)

	if email.Error != nil {
		log.Fatal(email.Error)
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
