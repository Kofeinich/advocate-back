package ui

import "advocate-back/internal/states"

func GetNewStateByActionID(curStateName string, id int, config states.BotStates) string {
	return config.States[curStateName].Actions[id].NextBlock
}
