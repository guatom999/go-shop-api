package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	"github.com/guatom999/go-shop-api/pkg/custom"
	_oauth2Exception "github.com/guatom999/go-shop-api/pkg/oauth2/exception"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {

	ctx := context.Background()

	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.playerTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsUserArePlayer(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.UnAuthorized{})
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.adminTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsUserAreAdmin(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.UnAuthorized{})
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {

	ctx := context.Background()

	updateToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.UnAuthorized{}
	}

	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, updateToken.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, updateToken.RefreshToken)

	return updateToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {

	ctx := context.Background()

	updateToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.UnAuthorized{}
	}

	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, updateToken.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, updateToken.RefreshToken)

	return updateToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(pctx echo.Context) (*oauth2.Token, error) {

	accessToken, err := pctx.Cookie(oauth2AccessTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.UnAuthorized{}
	}

	refreshToken, err := pctx.Cookie(oauth2RefreshTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.UnAuthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
