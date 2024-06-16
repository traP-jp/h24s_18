package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

type patchMeRequest struct {
	Bio string `json:"bio,omitempty"`
}

func PatchMe(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}
	data := &patchMeRequest{}
	err = c.Bind(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}
	model.UpdateUserBio(u.Name, data.Bio)
	return c.NoContent(http.StatusOK)
}

func GetMeHandler(c echo.Context) error {
	//ログインしてたらuserが返ってくる
	u, _, err := getMe(c)
	//もしエラーだったら
	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}
	//エラーじゃなかったら、データベースにユーザーの情報を問い合わせる
	user, err := model.GetUser(u.Name)
	//エラーだったら
	if err != nil {
		return err
	}
	//データベースから帰ってきたのを出力
	return c.JSON(http.StatusOK, user)
}
