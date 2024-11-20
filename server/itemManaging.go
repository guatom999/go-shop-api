package server

import (
	_itemManagingController "github.com/guatom999/go-shop-api/pkg/itemShop/controller"
	_itemManagingRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_itemManagingService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
)

func (s *echoServer) initItemManagingRouter() {
	itemManagingRepository := _itemManagingRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	_itemManagingService := _itemManagingService.NewItemShopServiceImpl(itemManagingRepository)
	_itemManagingController := _itemManagingController.NewItemShopControllerImpl(_itemManagingService)

	router := s.app.Group("/v1/item-managing")

	_ = router
	_ = _itemManagingController
}
