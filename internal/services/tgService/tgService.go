package tgService

import (
	"bot_forge_back/internal/delivery/http/validator"
	"bot_forge_back/internal/services/tgService/ui"
	"bot_forge_back/internal/states"
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

	updateType, err := CheckUpdateType(update)
	if err != nil {
		return SendMessageRequest{}, err
	}

	if updateType == MessageType {
		blockId, err := s.r.GetUserState(botId, strconv.Itoa(update.Message.From.ID))
		if err == redis.Nil {
			err = nil
			blockId = bot.InitialState
			s.r.SetUserState(botId, strconv.Itoa(update.Message.From.ID), blockId)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, bot.States[blockId].Text)
			msg.ReplyMarkup = ui.CreateTelegramKeyboard(blockId, bot)
			return SendMessageRequest{update, botId, botToken, msg, false}, err
		}
		newBlockId, err := ui.GetNewStateByActionText(blockId, bot)
		s.r.SetUserState(botId, strconv.Itoa(update.Message.From.ID), newBlockId)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, bot.States[newBlockId].Text)
		msg.ReplyMarkup = ui.CreateTelegramKeyboard(newBlockId, bot)
		return SendMessageRequest{update, botId, botToken, msg, false}, err
	}

	if updateType == CallbackQueryType {
		blockId, err := s.r.GetUserState(botId, strconv.Itoa(update.CallbackQuery.From.ID))
		if err != nil {
			return SendMessageRequest{}, nil
		}
		id, err := strconv.Atoi(update.CallbackQuery.Data)
		if err != nil {
			return SendMessageRequest{}, nil
		}
		newBlockId := ui.GetNewStateByActionID(blockId, id, bot)
		s.r.SetUserState(botId, strconv.Itoa(update.CallbackQuery.From.ID), newBlockId)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, bot.States[newBlockId].Text)
		msg.ReplyMarkup = ui.CreateTelegramKeyboard(newBlockId, bot)
		return SendMessageRequest{update, botId, botToken, msg, false}, err
	}
	return SendMessageRequest{}, nil
}
