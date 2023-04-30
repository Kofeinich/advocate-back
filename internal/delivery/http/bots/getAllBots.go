package bots

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h botHandler) GetAllBots(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, nil)
}
