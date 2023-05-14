package ui

import "bot_forge_back/internal/states"

func GetNewStateByActionID(curStateName string, id int, config states.BotStates) string {
	if len(config.States[curStateName].Actions) > id {
		return config.States[curStateName].Actions[id].NextBlock
	}
	return ""
}
