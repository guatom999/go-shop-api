package repository

import "github.com/guatom999/go-shop-api/entities"

type (
	PlayerCoinRepository interface {
		CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	}
)
