package controller

import (
	"net/http"

	"github.com/guatom999/go-shop-api/pkg/custom"
	_inventoryService "github.com/guatom999/go-shop-api/pkg/inventory/service"
	"github.com/guatom999/go-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type (
	inventoryControllerImpl struct {
		inventoryService _inventoryService.InventoryService
		logger           echo.Logger
	}
)

func NewInventoryControllerImpl(
	inventoryService _inventoryService.InventoryService,
	logger echo.Logger,
) InventoryController {
	return &inventoryControllerImpl{
		inventoryService,
		logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {

	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		c.logger.Errorf("error getting player id :%s", err.Error())
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)

}
