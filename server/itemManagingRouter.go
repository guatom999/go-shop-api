package server

import (
	_itemManagingController "github.com/guatom999/go-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/guatom999/go-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/guatom999/go-shop-api/pkg/itemManaging/service"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddleware) {

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemManagingRepository := _itemManagingRepository.NewItemManaginRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemShopRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router := s.app.Group("/v1/item-managing")

	router.POST("", itemManagingController.Creating, m.AdminAuthorizing)
	router.PATCH("/:itemID", itemManagingController.Editing, m.AdminAuthorizing)
	router.DELETE("/:itemID", itemManagingController.Archiving, m.AdminAuthorizing)

}
