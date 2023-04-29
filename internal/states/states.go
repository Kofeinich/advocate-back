package states

import "github.com/google/uuid"

type Button struct {
	Text      string
	NextBlock string
}

type State struct {
	Name    string
	Text    string
	Buttons []Button
	Alert   string
}

type BotStates struct {
	States       map[string]State
	CurrentState string
}

type Bot struct {
	Title     string
	Id        uuid.UUID
	BotStates *BotStates
}
