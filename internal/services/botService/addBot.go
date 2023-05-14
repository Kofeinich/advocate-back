package botService

import (
	"bot_forge_back/internal/services/tgService"
	"bot_forge_back/internal/states"
	"encoding/json"
	"github.com/google/uuid"
)

func (s Service) AddBot(conf states.BotStates, token string) error {
	botId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	hStr := botId.String()
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

	err = tgService.RegNewWebHook(token, botId.String())
	if err != nil {
		return err
	}
	return nil
}
