package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

func GetHello(c echo.Context) error {
	// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
	return c.String(http.StatusOK, "Hello, World.\n")
}

func GetUser(c echo.Context) error {
	userid := c.Param("userId")
	user, err := model.GetUser(userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
