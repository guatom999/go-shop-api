package service

import (
	_inventoryModel "github.com/guatom999/go-shop-api/pkg/inventory/model"
)

type (
	InventoryService interface {
		Listing(playerID string) ([]*_inventoryModel.Inventory, error)
	}
)
