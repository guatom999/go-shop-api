package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemManagingException "github.com/guatom999/go-shop-api/pkg/itemManaging/exception"
)

type (
	itemManagingRepositoryImpl struct {
		db     *gorm.DB
		logger echo.Logger
	}
)

func NewItemManaginRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {

	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed :%s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}
