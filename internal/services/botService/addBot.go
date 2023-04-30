package botService

import "github.com/google/uuid"

func (s Service) AddBot(conf string, token string) error {
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
