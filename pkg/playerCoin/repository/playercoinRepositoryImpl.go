package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	_playerCoinException "github.com/guatom999/go-shop-api/pkg/playerCoin/exception"
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

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {

	playerCoin := new(entities.PlayerCoin)

	if err := r.db.ConnectDatabase().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Errorf("player coin adding failed : %s", err.Error())
		return nil, &_playerCoinException.CoinAdding{}
	}

	return playerCoinEntity, nil
}
