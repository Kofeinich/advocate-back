package botService

import (
	"advocate-back/internal/states"
	"encoding/json"
)

func (s Service) UpdateBotConfig(conf states.BotStates, id string) error {
	marshalledJSON, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	err = s.r.CreateBotConfig(id, marshalledJSON)
	if err != nil {
		return err
	}
	return nil
}
