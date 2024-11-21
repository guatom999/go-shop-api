package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"

	_itemManagingException "github.com/guatom999/go-shop-api/pkg/itemManaging/exception"
	_itemManagingModel "github.com/guatom999/go-shop-api/pkg/itemManaging/model"
)

type (
	itemManagingRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewItemManaginRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {

	item := new(entities.Item)

	if err := r.db.ConnectDatabase().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed :%s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {

	if err := r.db.ConnectDatabase().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed :%s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archive(itemID uint64) error {

	if err := r.db.ConnectDatabase().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archive item failed :%s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
