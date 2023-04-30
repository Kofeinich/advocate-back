package states

const (
	ActionTypeButton ActionType = "button"
	ActionTypeText   ActionType = "text"
)

type ActionType string
type Action struct {
	Text      string
	NextBlock string
	Type      ActionType
}

type State struct {
	Name    string
	Text    string
	Actions []Action
	Alert   string
}

type BotStates struct {
	States       map[string]State
	CurrentState string
}

type Bot struct {
	Token     string
	BotStates *BotStates
}
