package service

import (
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
)

type (
	ItemShopService interface {
		Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
		Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
		Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
	}
)
