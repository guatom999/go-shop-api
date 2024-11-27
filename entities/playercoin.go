package entities

import (
	"time"

	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
)

type (
	PlayerCoin struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerId  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}
)

func (p *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
	return &_playerCoinModel.PlayerCoin{
		ID:        p.ID,
		PlayerID:  p.PlayerId,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}
