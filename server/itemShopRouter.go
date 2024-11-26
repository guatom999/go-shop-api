package server

import (
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopController "github.com/guatom999/go-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		inventoryRepository,
		playerCoinRepository,
		s.app.Logger,
	)
	_itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router := s.app.Group("/v1/item-shop")

	// _ = router

	router.GET("", _itemShopController.Listing)
	router.POST("", _itemShopController.Buying, m.PlayerAuthorizing)
}
