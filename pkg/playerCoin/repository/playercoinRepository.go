package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
)

type (
	PlayerCoinRepository interface {
		CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
		Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
	}
)
