package tgService

import (
	"advocate-back/internal/delivery/http/validator"
)

type Types string

const (
	MessageType            Types = "message"
	CallbackQueryType      Types = "callback_query"
	InlineQueryType        Types = "inline_query"
	ChosenInlineResultType Types = "chosen_inline_result"
	UnknownType            Types = "unknown"
)

func CheckUpdateType(update validator.TgValidatorRequest) Types {
	if update.Message != nil && update.Message.Text != "" {
		return MessageType
	} else if update.CallbackQuery != nil {
		return CallbackQueryType
	} else if update.InlineQuery != nil {
		return InlineQueryType
	} else if update.ChosenInlineResult != nil {
		return ChosenInlineResultType
	}

	return UnknownType
}
