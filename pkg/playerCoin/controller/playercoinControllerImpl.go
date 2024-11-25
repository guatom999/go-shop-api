package controller

import (
	"net/http"

	"github.com/guatom999/go-shop-api/pkg/custom"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
	_playerCoinService "github.com/guatom999/go-shop-api/pkg/playerCoin/service"
	"github.com/guatom999/go-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
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

func (c *playerCoinControllerImpl) CoinAdding(pctx echo.Context) error {

	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoReq := custom.NewCustomEchoRequest(pctx)

	if err := customEchoReq.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) Showing(pctx echo.Context) error {

	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoinShowing := c.playerCoinService.Showing(playerID)

	return pctx.JSON(http.StatusOK, playerCoinShowing)
}
