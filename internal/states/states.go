package states

const (
	ActionTypeButton ActionType = "button"
	ActionTypeText   ActionType = "text"
)

type ActionType string

func (a ActionType) IsValid() bool {
	for _, validType := range []ActionType{ActionTypeText, ActionTypeButton} {
		if a == validType {
			return true
		}
	}
	return false
}

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
	InitialState string
	States       map[string]State
}

type Bot struct {
	Token     string
	BotStates *BotStates
}
