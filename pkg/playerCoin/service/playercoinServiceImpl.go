package service

import (
	"github.com/guatom999/go-shop-api/entities"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
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

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {

	playerCoinEntity := &entities.PlayerCoin{
		PlayerId: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(nil, playerCoinEntity)
	if err != nil {
		return nil, err
	}

	playerCoinEntityResult.PlayerId = coinAddingReq.PlayerID

	return playerCoinEntityResult.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {

	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}

	return playerCoinShowing
}
