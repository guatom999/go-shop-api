package service

import (
	"github.com/guatom999/go-shop-api/entities"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
)

type (
	itemShopServiceImpl struct {
		itemShopRepository _itemShopRepository.ItemShopRepository
	}
)

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository: itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCounting, size)

	result := s.totalItemResponse(itemList, page, totalPage)

	return result, nil

}

func (s *itemShopServiceImpl) totalPageCalculation(totalItem int64, size int64) int64 {
	totalPage := totalItem / size

	if totalPage%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) totalItemResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {

	itemModelList := make([]*_itemShopModel.Item, 0)

	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
