package repository

import (
	"github.com/guatom999/go-shop-api/entities"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
)

type (
	ItemShopRepository interface {
		Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
		Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	}
)
