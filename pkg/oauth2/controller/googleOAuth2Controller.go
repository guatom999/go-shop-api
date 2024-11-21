package controller

import (
	"github.com/guatom999/go-shop-api/config"
	_oauth2Service "github.com/guatom999/go-shop-api/pkg/oauth2/service"
	"github.com/labstack/echo/v4"
)

type (
	googleOAuth2Controller struct {
		oauth2Service _oauth2Service.OAuth2Service
		oauth2Conf    *config.OAuth2
		logger        echo.Logger
	}
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) OAuth2Controller {
	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		logger:        logger,
	}
}
