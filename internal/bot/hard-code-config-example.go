package bot

import (
	"advocate-back/internal/states"
	"github.com/google/uuid"
)

var MockBotStates = states.BotStates{
	States: map[string]states.State{
		"start_message": {
			Name: "start_message",
			Text: "Привет! Я твой личный бот-ассистент. Что ты хочешь сделать?",
			Buttons: []states.Button{
				{
					Text:      "Отправить сообщение",
					NextBlock: "message_input",
				},
			},
			Alert: "Зачем тебе бот, если ты не знаешь, что делать?",
		},
		"message_input": {
			Name: "message_input",
			Text: "Введите ваше сообщение:",
			Buttons: []states.Button{
				{
					Text:      "Отправить",
					NextBlock: "send_message",
				},
			},
		},
	},
	CurrentState: "start_message",
}

var MockBot = states.Bot{
	Title:     "Mock Bot",
	Id:        uuid.New(),
	BotStates: &MockBotStates,
}
