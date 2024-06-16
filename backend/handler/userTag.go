package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/traP-jp/h24s_18/gemini"
	"github.com/traPtitech/go-traq"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

type PostTagRequest struct {
	TagName   string `json:"name" binding:"required"`
	IsStarred bool   `json:"isStarred" binding:"required"`
}

func PostUserTag(c echo.Context) error {
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

	err = insertUserTag(u, *body, c)

	if err != nil {
		return err
	}

	// エラーが起きなかったとき、正常なのでステータスコード 200 OK を返し、リクエストデータをそのまま返す
	return c.JSON(http.StatusOK, body)
}

func insertUserTag(u *traq.MyUserDetail, body PostTagRequest, c echo.Context) error {
	if body.TagName == "" {
		return c.String(http.StatusBadRequest, "tagname is blank") // http.StatusConflictがエラーの番号に該当する
	}

	err := model.CreateUserTag(u.Name, body.TagName, body.IsStarred)

	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return c.String(http.StatusConflict, "tag duplicate") // http.StatusConflictがエラーの番号に該当する
		}
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

func BulkInsertUserTags(c echo.Context) error {
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
		err = insertUserTag(u, tag, c)
		if err != nil {
			return err
		}
	}

	// エラーが起きなかったとき、正常なのでステータスコード 200 OK を返し、リクエストデータをそのまま返す
	return c.JSON(http.StatusOK, body)
}

func UpdateUserTag(c echo.Context) error {
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

func DeleteUserTag(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}

	tagName := c.Param("tagName")
	err = model.DeleteUserTag(u.Name, tagName) // サーバー起因のエラー
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return c.String(http.StatusOK, "delete "+tagName)
}
