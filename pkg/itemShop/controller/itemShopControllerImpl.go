package controller

import (
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
)

type (
	itemShopControllerImpl struct {
		itemShopService _itemShopService.ItemShopService
	}
)

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService: itemShopService}
}
