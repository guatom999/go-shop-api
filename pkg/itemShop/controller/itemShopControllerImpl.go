package controller

import (
	"net/http"

	"github.com/guatom999/go-shop-api/pkg/custom"
	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
	"github.com/guatom999/go-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type (
	itemShopControllerImpl struct {
		itemShopService _itemShopService.ItemShopService
	}
)

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService: itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {

	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, &_itemShopException.ItemListing{})
	}

	return pctx.JSON(http.StatusOK, itemModelList)

}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {

	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	buyingReq := new(_itemShopModel.BuyingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	buyingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, playerCoin)
}
