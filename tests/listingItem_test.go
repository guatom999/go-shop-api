package tests

import (
	"testing"

	"github.com/guatom999/go-shop-api/entities"
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"
	"github.com/stretchr/testify/assert"

	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"

	"github.com/labstack/echo/v4"
)

func TestListingItemSuccess(t *testing.T) {
	itemShopRepositoryMock := new(_itemShopRepository.ItemShopRepositoryMock)
	playerCoinRepositoryMock := new(_playerCoinRepository.PlayerCoinRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	echoLogger := echo.New().Logger

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
		echoLogger,
	)

	// _ = itemShopService

	itemShopRepositoryMock.On("Listing", &_itemShopModel.ItemFilter{
		Paginate: _itemShopModel.Paginate{
			Page: 1,
			Size: 10,
		},
	}).Return([]*entities.Item{
		{
			ID:          1,
			Name:        "Test",
			Description: "Sword Test",
			Picture:     "http://www.picture.com/test",
			Price:       100,
		},
		{
			ID:          2,
			Name:        "Test",
			Description: "Sword Test",
			Picture:     "http://www.picture.com/test",
			Price:       100,
		},
		{
			ID:          3,
			Name:        "Test",
			Description: "Sword Test",
			Picture:     "http://www.picture.com/test",
			Price:       100,
		},
	}, nil)

	itemShopRepositoryMock.On("Counting", &_itemShopModel.ItemFilter{
		Paginate: _itemShopModel.Paginate{
			Page: 1,
			Size: 10,
		},
	}).Return(
		int64(3),
		nil,
	)

	type args struct {
		label    string
		in       *_itemShopModel.ItemFilter
		expected *_itemShopModel.ItemResult
	}

	cases := []args{
		{
			label: "Success selling item",
			in: &_itemShopModel.ItemFilter{
				Paginate: _itemShopModel.Paginate{
					Page: 1,
					Size: 10,
				},
			},
			expected: &_itemShopModel.ItemResult{
				Items: []*_itemShopModel.Item{
					{
						ID:          1,
						Name:        "Test",
						Description: "Sword Test",
						Picture:     "http://www.picture.com/test",
						Price:       100,
					},
					{
						ID:          2,
						Name:        "Test",
						Description: "Sword Test",
						Picture:     "http://www.picture.com/test",
						Price:       100,
					},
					{
						ID:          3,
						Name:        "Test",
						Description: "Sword Test",
						Picture:     "http://www.picture.com/test",
						Price:       100,
					},
				},
				Paginate: _itemShopModel.PaginateResult{
					Page:      1,
					TotalPage: 0,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Listing(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)
		})
	}

}
