package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	logger echo.Logger
	db     databases.Database
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{logger, db}
}

func (r *itemShopRepositoryImpl) TransactionBegin() *gorm.DB {

	tx := r.db.ConnectDatabase()

	return tx.Begin()
}

func (r *itemShopRepositoryImpl) TransactionRollback(tx *gorm.DB) error {

	return tx.Rollback().Error
}

func (r *itemShopRepositoryImpl) TransactionCommit(tx *gorm.DB) error {

	return tx.Commit().Error
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {

	itemList := make([]*entities.Item, 0)

	query := r.db.ConnectDatabase().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil

}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {

	query := r.db.ConnectDatabase().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	count := new(int64)

	if err := query.Count(count).Error; err != nil {
		r.logger.Errorf("Counting items failed: %s", err.Error())
		return -1, &_itemShopException.ItemCounting{}
	}

	return *count, nil

}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.ConnectDatabase().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by id: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{ItemID: itemID}
	}

	return item, nil

}

func (r *itemShopRepositoryImpl) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	if err := r.db.ConnectDatabase().Model(&entities.Item{}).Where("id In ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID List: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return items, nil
}

func (r *itemShopRepositoryImpl) PurchaseHistoryRecording(tx *gorm.DB, purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {

	conn := r.db.ConnectDatabase()
	if tx != nil {
		conn = tx
	}

	insertedPurchasing := new(entities.PurchaseHistory)

	if err := conn.Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Failed to record purchase history: %s", err.Error())
		return nil, &_itemShopException.HistoryOfPurchaseRecording{}
	}

	return insertedPurchasing, nil
}
