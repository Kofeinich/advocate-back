package tgService

import (
	"advocate-back/internal/states"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

type Service struct {
	r userRepo
	b botRepo
}

type botRepo interface {
	GetBotConfigByID(botID string) (string, error)
}

type userRepo interface {
	SetUserState(botId string, userId string, blockId string) (err error)
	GetUserState(botId string, userId string) (blockId string, err error)
}

func NewService(r userRepo, b botRepo) *Service {
	return &Service{r: r, b: b}
}

func (s Service) ProcessTgUpdate(botId string, update tgbotapi.Update) (blockId string, err error) {
	botConfig, err := s.b.GetBotConfigByID(botId)
	bot := states.BotStates{}
	err = json.Unmarshal([]byte(botConfig), &bot)
	if err != nil {
		return "", err
	}
	updateType := CheckUpdateType(update)

	if updateType == FirstUpdate {

	}

	if updateType == MessageType {
		blockId, err = s.r.GetUserState(botId, strconv.Itoa(update.Message.From.ID))
		if err == redis.Nil {
			err = nil
			blockId = bot.InitialState
		}
		if err != nil {
			return "", err
		}
	}

	/// todo Message is possibly nil, check the type of event before getUserState

	/// todo switch analyze the update event check struct and change state

	/// todo for buttons data use cur block.Name + _ + arrayIndexOfAction

	/// todo prepare message config (MessageConfig) from newState

	/// todo when create messageConfig выставить ChatId, Text, ReplayMarkup

	// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/docs/examples/inline-keyboard.md

	return blockId, nil
}
