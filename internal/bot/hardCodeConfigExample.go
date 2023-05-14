package bot

import (
	"bot_forge_back/internal/states"
)

var MockBotStates = states.BotStates{
	InitialState: "start_message",
	States: map[string]states.State{
		"start_message": {
			Name: "start_message",
			Text: "Привет! Я твой личный бот-ассистент. Что ты хочешь сделать?",
			Actions: []states.Action{
				{
					Text:      "Отправить сообщение",
					NextBlock: "message_input",
					Type:      states.ActionTypeButton,
				},
			},
			Alert: "Зачем тебе бот, если ты не знаешь, что делать?",
		},
		"message_input": {
			Name: "message_input",
			Text: "Введите ваше сообщение:",
			Actions: []states.Action{
				{
					NextBlock: "message_sent",
					Type:      states.ActionTypeText,
				},
				{
					Text:      "Назад",
					NextBlock: "start_message",
					Type:      states.ActionTypeButton,
				},
			},
		},
		"message_sent": {
			Name: "message_sent",
			Text: "Сообщение отправлено",
			Actions: []states.Action{
				{
					Text:      "Отправить заново",
					NextBlock: "start_message",
					Type:      "button",
				},
				{
					Text:      "Назад",
					NextBlock: "start_message",
					Type:      "button",
				},
			},
		},
	},
}

var MockBot = states.Bot{
	Token:     "ffodkfofk",
	BotStates: &MockBotStates,
}
