package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/guatom999/go-shop-api/config"
	"github.com/guatom999/go-shop-api/databases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_adminRepository "github.com/guatom999/go-shop-api/pkg/admin/repository"
	_oauth2Controller "github.com/guatom999/go-shop-api/pkg/oauth2/controller"
	_oauth2Service "github.com/guatom999/go-shop-api/pkg/oauth2/service"
	_playerRepository "github.com/guatom999/go-shop-api/pkg/player/repository"
)

type (
	echoServer struct {
		app  *echo.Echo
		db   databases.Database
		conf *config.Config
	}
)

var (
	server *echoServer
	once   sync.Once
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *echoServer) Start() {

	timeoutMiddleware := getTimeOutMiddleware(s.conf.Server.TimeOut)
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowedOrigins)
	bodyLimitMiddleware := getBodyLimitMiddlware(s.conf.Server.BodyLimit)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeoutMiddleware)

	authorizingMiddleware := s.getAuthorizingMiddleware()

	s.app.GET("/v1/health", s.healthCheck)

	s.initItemShopRouter(authorizingMiddleware)
	s.initItemManagingRouter(authorizingMiddleware)
	s.initOAuth2Router()
	s.initPlayerCoinRouter(authorizingMiddleware)
	s.initInventoryRouter(authorizingMiddleware)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	go s.gratefullyShutdown(quitCh)
	s.httpListening()
}

func (s *echoServer) httpListening() {
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(serverUrl); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) gratefullyShutdown(quitCh chan os.Signal) {

	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("Shutting down...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout:      timeout * time.Second,
	})
}

func getCORSMiddleware(allowOrigin []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigin,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddlware(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}

func (s *echoServer) getAuthorizingMiddleware() *authorizingMiddleware {
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuthService(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		s.app.Logger,
	)

	return &authorizingMiddleware{
		oauth2Controller: oauth2Controller,
		oauth2Conf:       s.conf.OAuth2,
		logger:           s.app.Logger,
	}
}
