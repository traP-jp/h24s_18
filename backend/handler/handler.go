package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetHello(c echo.Context) error {
	// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
	return c.String(http.StatusOK, "Hello, World.\n")
}
