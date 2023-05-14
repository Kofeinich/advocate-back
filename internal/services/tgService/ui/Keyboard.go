package ui

import (
	"bot_forge_back/internal/states"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func CreateTelegramKeyboard(curStateName string, config states.BotStates) tgbotapi.InlineKeyboardMarkup {
	if _, ok := config.States[curStateName]; !ok || len(config.States[curStateName].Actions) == 0 {
		return tgbotapi.InlineKeyboardMarkup{}
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for index, button := range config.States[curStateName].Actions {
		inlineButton := tgbotapi.NewInlineKeyboardButtonData(button.Text, strconv.Itoa(index))
		row := []tgbotapi.InlineKeyboardButton{inlineButton}
		rows = append(rows, row)
	}
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}
