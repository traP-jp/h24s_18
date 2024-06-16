package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
	"net/http"
)

func GetUser(c echo.Context) error {
	userid := c.Param("userId")
	user, err := model.GetUser(userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
