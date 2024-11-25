package service

import (
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
)

type (
	PlayerCoinService interface {
		CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	}
)
