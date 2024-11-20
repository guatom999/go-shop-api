package controller

import (
	"net/http"

	"github.com/guatom999/go-shop-api/pkg/custom"
	_itemManagingModel "github.com/guatom999/go-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/guatom999/go-shop-api/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type (
	itemManagingControllerImpl struct {
		itemManagingService _itemManagingService.ItemManagingService
	}
)

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Createing(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, item)
}
