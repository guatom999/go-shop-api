package repository

import "github.com/guatom999/go-shop-api/entities"

type (
	InventoryRepository interface {
		Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
		Removing(playerID string, itemID uint64, limit int) error
		PlayerItemCounting(playerID string, itemID uint64) int64
		Listing(playerID string) ([]*entities.Inventory, error)
	}
)
