package tgService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Types string

const (
	MessageType            Types = "message"
	CallbackQueryType      Types = "callback_query"
	InlineQueryType        Types = "inline_query"
	ChosenInlineResultType Types = "chosen_inline_result"
	FirstUpdate            Types = "first_update"
)

func CheckUpdateType(update tgbotapi.Update) Types {
	if update.Message != nil {
		return MessageType
	} else if update.CallbackQuery != nil {
		return CallbackQueryType
	} else if update.InlineQuery != nil {
		return InlineQueryType
	} else if update.ChosenInlineResult != nil {
		return ChosenInlineResultType
	}
	if update.Message == nil || update.Message.Text == "" {
		return FirstUpdate
	}
	return ""
}
