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
	Text      string     `json:"text"`
	NextBlock string     `json:"next_block,omitempty"`
	Type      ActionType `json:"type"`
}

type State struct {
	Name    string   `json:"name"`
	Text    string   `json:"text"`
	Actions []Action `json:"actions"`
	Alert   string   `json:"alert,omitempty"`
}

type BotStates struct {
	InitialState string           `json:"initial_state"`
	States       map[string]State `json:"states"`
}

type Bot struct {
	Token     string     `json:"tg_token"`
	BotStates *BotStates `json:"bot_config"`
}
