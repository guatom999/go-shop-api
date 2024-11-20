package server

import (
	_itemManagingController "github.com/guatom999/go-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/guatom999/go-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/guatom999/go-shop-api/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	itemManagingRepository := _itemManagingRepository.NewItemManaginRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router := s.app.Group("/v1/item-managing")

	router.POST("", itemManagingController.Creating)

}
