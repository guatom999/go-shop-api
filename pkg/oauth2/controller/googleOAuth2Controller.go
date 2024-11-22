package controller

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/guatom999/go-shop-api/config"
	"github.com/guatom999/go-shop-api/pkg/custom"
	_oauth2Exception "github.com/guatom999/go-shop-api/pkg/oauth2/exception"
	_oauth2Model "github.com/guatom999/go-shop-api/pkg/oauth2/model"
	_oauth2Service "github.com/guatom999/go-shop-api/pkg/oauth2/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type (
	googleOAuth2Controller struct {
		oauth2Service _oauth2Service.OAuth2Service
		oauth2Conf    *config.OAuth2
		logger        echo.Logger
	}
)

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	oauth2AccessTokenCookieName  = "act"
	oauth2RefreshTokenCookieName = "rft"
	stateCookieName              = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) OAuth2Controller {

	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		logger:        logger,
	}
}

func setGoogleOAuth2Config(oauth2Config *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Config.ClientId,
		ClientSecret: oauth2Config.ClientSecret,
		RedirectURL:  oauth2Config.PlayerRedirectUrl,
		Scopes:       oauth2Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Config.Endpoints.AuthUrl,
			TokenURL:      oauth2Config.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Config.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Config.ClientId,
		ClientSecret: oauth2Config.ClientSecret,
		RedirectURL:  oauth2Config.AdminRedirectUrl,
		Scopes:       oauth2Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Config.Endpoints.AuthUrl,
			TokenURL:      oauth2Config.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Config.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	state := c.randomState()

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	state := c.randomState()

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {

	// ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Error("Failed to validate callback: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err.Error())
	}

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "login successful"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {

	// ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Error("Failed to validate callback: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err.Error())
	}

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "login successful"})
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {

	c.removeCookie(pctx, oauth2AccessTokenCookieName)
	c.removeCookie(pctx, oauth2RefreshTokenCookieName)
	c.removeCookie(pctx, stateCookieName)

	return pctx.NoContent(http.StatusNoContent)
}

func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) {

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		// Secure: true, use on production
	}

	pctx.SetCookie(cookie)

}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		// Secure: true, use on production
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) callbackValidating(pctx echo.Context) error {

	state := pctx.QueryParam("state")

	stateFromCookie, err := pctx.Cookie(stateCookieName)
	if err != nil {
		c.logger.Errorf("Failed to get state from cookies :%s", err.Error())
		return &_oauth2Exception.UnAuthorized{}
	}

	if state == "" || state != stateFromCookie.Value {
		c.logger.Errorf("Invalid state: != %s", state)
		return &_oauth2Exception.UnAuthorized{}
	}

	return nil

}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
