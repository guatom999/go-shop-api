package model

import "time"

type (
	PlayerCoin struct {
		ID        string    `json:"id"`
		PlayerID  string    `json:"playerID"`
		Amount    int64     `json:"amount"`
		CreatedAT time.Time `json:"createdAT"`
	}

	CoinAddingReq struct {
		PlayerID string
		Amount   int64 `json:"amount" validate="required,gt=0"`
	}

	PlayerCoinShowing struct {
		PlayerID string `json:"playerID"`
		Coin     int64  `json:"coin"`
	}
)
