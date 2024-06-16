package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
)

type GetUserResponse struct {
	Name          string
	Id            string
	Bio           string
	TwitterId     string
	HomeChannelId string
	Tag           []model.UserTag
}

func GetUser(c echo.Context) error {
	userid := c.Param("userId")
	user, err := model.GetUser(userid)
	if err != nil {
		return err
	}

	tags, err := model.GetUserTagsByUserId(userid)
	if err != nil {
		return err
	}

	response := GetUserResponse{
		Name:          user.Name,
		Id:            user.Id,
		Bio:           user.Bio,
		TwitterId:     user.TwitterId,
		HomeChannelId: user.HomeChannelId,
		Tag:           tags,
	}
	return c.JSON(http.StatusOK, response)
}
