package controller

import (
	"net/http"
	"strconv"

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
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Createing(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingControllerImpl) Editing(pctx echo.Context) error {

	// itemID := pctx.Param("itemID")
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemEditReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemEditReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Editing(itemID, itemEditReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingControllerImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)

	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusOK)
}

func (c *itemManagingControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}
	return itemIDUint64, nil
}
