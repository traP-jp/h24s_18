package handler

import (
	"errors"
	"fmt"
	"github.com/traP-jp/h24s_18/gemini"
	"github.com/traPtitech/go-traq"
	"gorm.io/gorm"
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

	err = insertTag(u, *body, c)

	if err != nil {
		return err
	}

	// エラーが起きなかったとき、正常なのでステータスコード 200 OK を返し、リクエストデータをそのまま返す
	return c.JSON(http.StatusOK, body)
}

func insertTag(u *traq.MyUserDetail, body PostTagRequest, c echo.Context) error {
	err := model.CreateUserTag(u.Name, body.TagName, body.IsStarred)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	_, err = model.GetTag(body.TagName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			value, err := gemini.GetEmbedding(body.TagName, c.Request().Context())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
			}
			err = model.CreateTag(body.TagName, value.String())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
			}
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}
	} else {
		// do nothing
	}

	return nil
}

type BulkInsertTagsRequest []PostTagRequest

func BulkInsertTags(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}

	// 受け取りたい JSON を示す空の変数を先に用意する。
	body := &BulkInsertTagsRequest{}
	// 受け取った JSON を data に代入する
	err = c.Bind(body)

	if err != nil { // エラーが発生した時、以下を実行
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	for _, tag := range *body {
		err = insertTag(u, tag, c)
		if err != nil {
			return err
		}
	}

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
