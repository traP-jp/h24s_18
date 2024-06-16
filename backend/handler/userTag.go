package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

type PostTagRequest struct {
	TagName   string `json:"name" binding:"required"`
	IsStarred bool   `json:"isStarred" binding:"required"`
}

func PostTag(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}

	// 受け取りたい JSON を示す空の変数を先に用意する。
	body := &PostTagRequest{}
	// 受け取った JSON を data に代入する
	err = c.Bind(body)

	if err != nil { // エラーが発生した時、以下を実行
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	model.CreateUserTag(u.Name, body.TagName, body.IsStarred)

	// エラーが起きなかったとき、正常なのでステータスコード 200 OK を返し、リクエストデータをそのまま返す
	return c.JSON(http.StatusOK, body)
}

func UpdateTag(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}

	// 受け取りたい JSON を示す空の変数を先に用意する。
	body := &PostTagRequest{}
	// 受け取った JSON を data に代入する
	err = c.Bind(body)
	if err != nil { // エラーが発生した時、以下を実行 リクエスト起因のエラー
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}
	err = model.UpdateUserTags(u.Name, body.TagName, body.IsStarred) // サーバー起因のエラー
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return c.JSON(http.StatusOK, body)
}

type DeleteTagRequest struct {
	TagName string `json:"name" binding:"required"`
}

func DeleteTag(c echo.Context) error {
	// 受け取りたい JSON を示す空の変数を先に用意する。
	body := &DeleteTagRequest{}
	// 受け取った JSON を data に代入する
	err := c.Bind(body)
	if err != nil { // エラーが発生した時、以下を実行 リクエスト起因のエラー
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}
	err = model.DeleteUserTag(body.TagName) // サーバー起因のエラー
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return c.JSON(http.StatusOK, body)
}
