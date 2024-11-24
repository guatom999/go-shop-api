package controller

import (
	_playerCoinService "github.com/guatom999/go-shop-api/pkg/playerCoin/service"
)

type (
	playerCoinControllerImpl struct {
		playerCoinService _playerCoinService.PlayerCoinService
	}
)

func NewPlayerCoinControllerImpl(
	playerCoinService _playerCoinService.PlayerCoinService,
) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}
