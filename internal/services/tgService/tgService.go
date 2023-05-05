package tgService

import (
	"advocate-back/internal/delivery/http/validator"
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
	GetBotTokenByID(botID string) (string, error)
}

type userRepo interface {
	SetUserState(botId string, userId string, blockId string) (err error)
	GetUserState(botId string, userId string) (blockId string, err error)
}

func NewService(r userRepo, b botRepo) *Service {
	return &Service{r: r, b: b}
}
func (s Service) ProcessTgUpdate(botId string, update validator.TgValidatorRequest) (sendRequest SendMessageRequest, err error) {
	botConfig, err := s.b.GetBotConfigByID(botId)
	botToken, err := s.b.GetBotTokenByID(botId)
	if err != nil {
		return SendMessageRequest{}, err
	}

	bot := states.BotStates{}
	err = json.Unmarshal([]byte(botConfig), &bot)
	if err != nil {
		return SendMessageRequest{}, err
	}

	updateType := CheckUpdateType(update)

	if updateType == MessageType {
		blockId, err := s.r.GetUserState(botId, strconv.Itoa(update.Message.From.ID))
		if err == redis.Nil {
			err = nil
			blockId = bot.InitialState
			s.r.SetUserState(botId, strconv.Itoa(update.Message.From.ID), blockId)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
			return SendMessageRequest{update, botId, botToken, msg}, err
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
		return SendMessageRequest{update, botId, botToken, msg}, err

		// TODO: Analyze the update event and change state accordingly.

		// TODO: For button data, use the current block.Name + _ + the index of the action.

		// TODO: Prepare message config (MessageConfig) from newState.

		// TODO: When creating MessageConfig, set ChatID, Text, and ReplyMarkup.

		// See: https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/docs/examples/inline-keyboard.md.

	}

	return SendMessageRequest{}, nil
}
