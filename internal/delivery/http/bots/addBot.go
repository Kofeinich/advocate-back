package bots

import (
	"bot_forge_back/internal/algorithm"
	validate "bot_forge_back/internal/delivery/http/validator"
	"bot_forge_back/internal/states"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h BotHandler) AddBot(c echo.Context) (err error) {
	m := new(validate.AddBotRequest)
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// todo rewrite into service

	botConfigBytes, err := json.Marshal(m.BotConfig)
	if err != nil {
		return err
	}

	var botStates states.BotStates
	if err := json.Unmarshal(botConfigBytes, &botStates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = algorithm.CheckAlgorithm(botStates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.s.AddBot(m.BotConfig, m.Token); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}
