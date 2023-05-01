package algorithm

import (
	"advocate-back/internal/states"
	"fmt"
)

func ValidateReferences(s *states.BotStates) error {
	stateMap := make(map[string]states.State)
	for _, state := range s.States {
		stateMap[state.Name] = state
	}

	if _, ok := stateMap[s.InitialState]; !ok {
		return fmt.Errorf("Initial state '%s' not found in states", s.InitialState)
	}

	for key, state := range s.States {

		if state.Name != key {
			return fmt.Errorf("Error in '%s' block name is not equal to key:'%s'", state, state.Name)
		}

		for _, button := range state.Actions {
			if !button.Type.IsValid() {
				return fmt.Errorf("Error in '%s' Action Type is not valid in '%s'", button, state.Name)
			}
			if _, ok := stateMap[button.NextBlock]; !ok {
				return fmt.Errorf("NextBlock '%s' in state '%s' does not exist", button.NextBlock, state.Name)
			}
		}
	}
	return nil
}
