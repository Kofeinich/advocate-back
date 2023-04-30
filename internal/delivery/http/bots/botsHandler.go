package bots

type botHandler struct {
	s service
}

type service interface {
}

func NewBotHandler(s service) *botHandler {
	return &botHandler{s: s}
}
