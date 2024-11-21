package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	logger echo.Logger
	db     *gorm.DB
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{logger, db}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {

	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

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

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

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

	if err := r.db.First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by id: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{ItemID: itemID}
	}

	return item, nil

}
