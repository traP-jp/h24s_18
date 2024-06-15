package handler

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
	traqoauth2 "github.com/traPtitech/go-traq-oauth2"
	"golang.org/x/oauth2"
)

const (
	sessionName = "session_name"
	stateLength = 16
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     os.Getenv("TRAQ_CLIENT_ID"),
		ClientSecret: os.Getenv("TRAQ_CLIENT_SECRET"),
		Endpoint:     traqoauth2.Prod,
		RedirectURL:  os.Getenv("TRAQ_REDIRECT_URL"), // e.g. http://localhost:8080/oauth2/callback,
	}

	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)

func GetMeHandler(c echo.Context) error {
	u, _, err := getMe(c)

	if err != nil {
		if errors.Is(err, errUnauthorized) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return err
	}
	user, err := model.GetUser(u.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func AuthorizeHandler(c echo.Context) error {
	session, err := store.Get(c.Request(), sessionName)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	codeVerifier := oauth2.GenerateVerifier()
	session.Values["code_verifier"] = codeVerifier

	state := make([]byte, stateLength)
	_, _ = rand.Read(state)
	session.Values["state"] = string(state)

	if err := session.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	// this client forces to use PKCE
	// code_challenge_method = S256 is set by S256ChallengeOption
	authCodeURL := oauth2Config.AuthCodeURL(
		string(state),
		oauth2.S256ChallengeOption(codeVerifier),
	)

	http.Redirect(c.Response(), c.Request(), authCodeURL, http.StatusFound)
	return nil
}

func CallbackHandler(c echo.Context) error {
	// query parameters
	var (
		code  = c.QueryParam("code")
		state = c.QueryParam("state")
	)

	if code == "" {
		return c.String(http.StatusBadRequest, "code is required")
	}

	session, err := store.Get(c.Request(), sessionName)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	codeVerifier, ok := session.Values["code_verifier"]
	if !ok {
		return c.String(http.StatusBadRequest, "code_verifier is required")
	}

	if storedState, ok := session.Values["state"]; !ok || storedState.(string) != state {
		return c.String(http.StatusBadRequest, "state is invalid")
	}

	token, err := oauth2Config.Exchange(
		c.Request().Context(),
		code,
		oauth2.VerifierOption(codeVerifier.(string)),
	)
	if err != nil {
		return fmt.Errorf("failed to exchange token: %w", err)
	}

	session.Values["token"] = token
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	// Save data to database

	return c.String(http.StatusOK, "success")
}
