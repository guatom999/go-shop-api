package server

import (
	_itemShopController "github.com/guatom999/go-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	_itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router := s.app.Group("/v1/item-shop")

	// _ = router

	router.GET("", _itemShopController.Listing)
}
