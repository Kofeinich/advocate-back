package bots

import (
	"advocate-back/internal/algorithm"
	validate "advocate-back/internal/delivery/http/validator"
	"advocate-back/internal/states"
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

	botConfigBytes, err := json.Marshal(m.BotConfig)
	if err != nil {
		return err
	}

	var botStates states.BotStates
	if err := json.Unmarshal(botConfigBytes, &botStates); err != nil {
		return err
	}

	if err = algorithm.CheckAlgorithm(botStates); err != nil {
		return err
	}

	if err = h.s.AddBot(m.BotConfig, m.Token); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}
