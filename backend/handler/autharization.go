package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/go-traq"
	"golang.org/x/oauth2"
	"net/http"
)

var (
	errUnauthorized = fmt.Errorf("unauthorized")
)

func getMe(c echo.Context) (*traq.MyUserDetail, *traq.APIClient, error) {
	session, err := store.Get(c.Request(), sessionName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get session: %w", err)
	}

	token, ok := session.Values["token"].(*oauth2.Token)
	if !ok {
		return nil, nil, c.String(http.StatusBadRequest, "token is required")
	}

	ctx := c.Request().Context()

	traqConfig := traq.NewConfiguration()
	traqConfig.HTTPClient = oauth2Config.Client(ctx, token)

	cli := traq.NewAPIClient(traqConfig)

	user, res, err := cli.MeApi.GetMe(ctx).Execute()
	if err != nil {
		// unauthorized
		return nil, nil, err
	}

	if res.StatusCode != 200 {
		fmt.Printf("failed to get me: %d\n", res.StatusCode)
		return nil, nil, errUnauthorized
	}

	return user, traq.NewAPIClient(traqConfig), nil
}
