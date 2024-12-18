package service

import (
	_itemManagingModel "github.com/guatom999/go-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
)

type (
	ItemManagingService interface {
		Createing(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
		Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
		Archiving(itemID uint64) error
	}
)
