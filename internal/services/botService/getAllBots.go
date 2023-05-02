package botService

import (
	"advocate-back/internal/states"
	"encoding/json"
)

type Bot struct {
	Id    string           `json:"id"`
	Conf  states.BotStates `json:"conf"`
	Token string           `json:"token"`
}

func (s Service) GetAllBots() ([]Bot, error) {
	bots, err := s.r.GelAllBotsFromList()
	if err != nil {
		return nil, err
	}
	botsArr := make([]Bot, 0, len(bots))
	for _, botID := range bots {
		conf, err := s.r.GetBotConfigByID(botID)
		if err != nil {
			return nil, err
		}
		var statesMap states.BotStates
		err = json.Unmarshal([]byte(conf), &statesMap)
		if err != nil {
			return nil, err
		}
		token, err := s.r.GetBotTokenByID(botID)
		if err != nil {
			return nil, err
		}
		bot := Bot{
			Id:    botID,
			Conf:  statesMap,
			Token: token,
		}
		botsArr = append(botsArr, bot)

	}
	return botsArr, nil
}
