package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	_itemManagingModel "github.com/guatom999/go-shop-api/pkg/itemManaging/model"
)

type (
	ItemManagingRepository interface {
		Creating(itemEntity *entities.Item) (*entities.Item, error)
		Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
		Archive(itemId uint64) error
	}
)
