package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
