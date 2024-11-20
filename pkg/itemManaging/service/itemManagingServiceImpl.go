package service

import (
	_itemManagingRepository "github.com/guatom999/go-shop-api/pkg/itemManaging/repository"
)

type (
	itemManagingServiceImpl struct {
		itemManagingRepository _itemManagingRepository.ItemManagingRepository
	}
)

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}
