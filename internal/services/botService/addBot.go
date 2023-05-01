package botService

import (
	"advocate-back/internal/states"
	"encoding/json"
	"github.com/google/uuid"
)

func (s Service) AddBot(conf states.BotStates, token string) error {
	n, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	hStr := n.String()
	marshalledJSON, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	err = s.r.AddBotToBotsList(hStr)
	if err != nil {
		return err
	}
	err = s.r.CreateBotToken(hStr, token)
	if err != nil {
		return err
	}
	err = s.r.CreateBotConfig(hStr, marshalledJSON)
	if err != nil {
		return err
	}
	return nil
}
