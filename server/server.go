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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type (
	echoServer struct {
		app  *echo.Echo
		db   *gorm.DB
		conf *config.Config
	}
)

var (
	server *echoServer
	once   sync.Once
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
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

	s.app.GET("/v1/health", s.healthCheck)

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
