package ui

import (
	"bot_forge_back/internal/states"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNewStateByActionText(curStateName string, config states.BotStates) (string, error) {
	for _, action := range config.States[curStateName].Actions {
		if action.Type == states.ActionTypeText {
			return action.NextBlock, nil
		}
	}
	return "", echo.NewHTTPError(http.StatusBadRequest)
}
