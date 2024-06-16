package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

func GetTags(c echo.Context) error {
	tagname, err := model.GetAllTagsName()
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, tagname)

}

func FindUserByTag(c echo.Context) error {
	q := c.QueryParam("q")
	user, err := model.RecommendUserByTag(q)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
