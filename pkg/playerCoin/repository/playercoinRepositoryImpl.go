package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/labstack/echo/v4"
)

type (
	playerCoinRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewPlayerCoinRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db, logger}
}
