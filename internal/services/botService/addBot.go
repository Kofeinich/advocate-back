package botService

import "github.com/google/uuid"

type service struct {
	r repo
}

type repo interface {
	DeleteBotFromList(botID string) error
	GelAllBotsFromList() error
	AddBotToBotsList(botID string) error
	CreateBotConfig(botID string, config string) error
	GetBotConfigByID(botID string) (string, error)
	CreateBotToken(botID string, token string) error
	GetBotTokenByID(botID string) (string, error)
}

func NewService(r repo) *service {
	return &service{r: r}

}

func (s service) AddBot(conf string, token string) error {
	n, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	hStr := n.String()
	err = s.r.AddBotToBotsList(hStr)
	if err != nil {
		return err
	}
	err = s.r.CreateBotToken(hStr, token)
	if err != nil {
		return err
	}
	err = s.r.CreateBotConfig(hStr, conf)
	if err != nil {
		return err
	}
	return nil
}
