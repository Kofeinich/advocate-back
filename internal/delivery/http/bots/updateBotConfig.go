package bots

import (
	"bot_forge_back/internal/algorithm"
	validate "bot_forge_back/internal/delivery/http/validator"
	"bot_forge_back/internal/states"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h BotHandler) UpdateBotConfig(c echo.Context) (err error) {
	m := new(validate.UpdateBotConfigRequest)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	botConfigBytes, err := json.Marshal(m.BotConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var botStates states.BotStates
	if err = json.Unmarshal(botConfigBytes, &botStates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = algorithm.CheckAlgorithm(botStates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.s.UpdateBotConfig(m.BotConfig, m.BotID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, m)
}
