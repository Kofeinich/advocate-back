package bots

type BotHandler struct {
	s service
}

type service interface {
	AddBot(conf string, token string) error
	DeleteBot(id string) error
	GetAllBots() ([]string, error)
	UpdateBotConfig(conf string, id string) error
}

func NewBotHandler(s service) *BotHandler {
	return &BotHandler{s: s}
}
