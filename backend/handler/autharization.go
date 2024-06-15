package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/go-traq"
	"golang.org/x/oauth2"
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
		fmt.Println("token is required")
		return nil, nil, errUnauthorized
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

	fmt.Println(res.StatusCode)

	if res.StatusCode != 200 {
		fmt.Printf("failed to get me: %d\n", res.StatusCode)
		return nil, nil, errUnauthorized
	}

	return user, traq.NewAPIClient(traqConfig), nil
}
