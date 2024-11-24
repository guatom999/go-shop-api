package service

import (
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"
)

type (
	playerCoinServiceImpl struct {
		playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	}
)

func NewPlayerCoinServiceImpl(
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}
