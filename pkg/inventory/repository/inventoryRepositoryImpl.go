package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_inventoryException "github.com/guatom999/go-shop-api/pkg/inventory/exception"
)

type (
	inventoryRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewInventoryRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) InventoryRepository {
	return &inventoryRepositoryImpl{db, logger}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {

	conn := r.db.ConnectDatabase()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for i := 0; i < qty; i++ {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("error filling inventory: %s", err)
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {

	conn := r.db.ConnectDatabase()
	if tx != nil {
		conn = tx
	}

	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := conn.Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory : %s", err)
			return &_inventoryException.PlayerItemRemoving{ItemID: inventory.ID}
		}
	}

	// if err := tx.Commit().Error; err != nil {
	// 	r.logger.Errorf("error removing player item in inventory : %s", err)
	// 	return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
	// }

	return nil
}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {

	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.ConnectDatabase().Where(
		"player_id = ? and item_id = ? and is_deleted = ?",
		playerID,
		itemID,
		false,
	).Limit(
		limit,
	).Find(
		inventoryEntities,
	).Error; err != nil {
		r.logger.Errorf("error finding player item in inventory by ID : %s", err)
		return nil, &_inventoryException.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventoryEntities, nil

}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {

	var count int64

	if err := r.db.ConnectDatabase().Model(&entities.Inventory{}).Where(
		"player_id = ? and item_id = ? and is_deleted = ?",
		playerID,
		itemID,
		false,
	).Count(&count).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory : %s", err)
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {

	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.ConnectDatabase().Where("player_id = ? and is_deleted = ?", playerID, false).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory : %s", err)
		return nil, &_inventoryException.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventoryEntities, nil
}
