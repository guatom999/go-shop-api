package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	_playerCoinException "github.com/guatom999/go-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db,
		logger,
	}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {

	playerCoin := new(entities.PlayerCoin)

	if err := r.db.ConnectDatabase().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Errorf("player coin adding failed : %s", err.Error())
		return nil, &_playerCoinException.CoinAdding{}
	}

	return playerCoinEntity, nil
}

func (r *playerCoinRepositoryImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {

	playerCoinShowing := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.ConnectDatabase().Model(&entities.PlayerCoin{}).Where("player_id = ?", playerID).Select("player_id , sum(amount) as coin").Group("player_id").Scan(playerCoinShowing).Error; err != nil {
		r.logger.Errorf("player coin adding failed : %s", err.Error())
		return nil, &_playerCoinException.PlayerCoinShowing{}
	}

	return playerCoinShowing, nil

}
