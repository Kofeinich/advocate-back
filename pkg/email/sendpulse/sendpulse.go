package sendpulse

import (
	sp "github.com/ont/go-sendpulse"
	"log"
	"time"
)

func SendPulseMessage() {
	var d time.Duration = 1000000000
	s := sp.New("39d5b0d07ad41f6bbbd7fd45464f207e", "5d3e28547d47ec018ddae275463ee2b4", d, true)

	email, err := sp.NewEmail(
		sp.Address{Name: "from me", Email: "sender@some.com"},
		sp.Address{Name: "to this person", Email: "kvr165@mail.ru"},
		"test subject",
		"<b>test</b> html",
		"test text",
	)
	log.Println(err)
	err = s.SMTP.Send(email)
	log.Println(err)
}
