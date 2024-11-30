package tests

import (
	"testing"

	"github.com/guatom999/go-shop-api/entities"
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"

	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestItemSellingSuccess(t *testing.T) {
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

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("TransactionBegin").Return(tx)
	itemShopRepositoryMock.On("TransactionRollback", tx).Return(nil)
	itemShopRepositoryMock.On("TransactionCommit", tx).Return(nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
		Price:       1000,
	}, nil)

	// inventoryRepositoryMock.On("PlayerItemCounting", "P0001", uint64(1)).Return(int64(2))
	inventoryRepositoryMock.On("PlayerItemCounting", "P0001", uint64(1)).Return(int64(2))

	playerCoinRepositoryMock.On("Showing", "P0001").Return(&_playerCoinModel.PlayerCoinShowing{
		PlayerID: "P0001",
		Coin:     2000,
	},
		nil)

	itemShopRepositoryMock.On("PurchaseHistoryRecording", &entities.PurchaseHistory{
		PlayerID:        "P0001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPrice:       1000,
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		IsBuying:        false,
		Quantity:        2,
	}, tx).Return(&entities.PurchaseHistory{
		PlayerID:        "P0001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPrice:       1000,
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		IsBuying:        false,
		Quantity:        2,
	}, nil)

	playerCoinRepositoryMock.On("CoinAdding", tx, &entities.PlayerCoin{
		PlayerId: "P0001",
		Amount:   1000,
	}).Return(&entities.PlayerCoin{
		PlayerId: "P0001",
		Amount:   1000,
	}, nil)

	inventoryRepositoryMock.On("Removing", tx, "P0001", uint64(1), 2).Return(nil)

	type args struct {
		label    string
		in       *_itemShopModel.SellingReq
		expected *_playerCoinModel.PlayerCoin
	}

	cases := []args{
		{
			label: "Success selling item",
			in: &_itemShopModel.SellingReq{
				PlayerID: "P0001",
				ItemID:   1,
				Quantity: 2,
			},
			expected: &_playerCoinModel.PlayerCoin{
				PlayerID: "P0001",
				Amount:   1000,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Selling(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)
		})
	}

}

func TestItemSellingFailed(t *testing.T) {
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

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("BeginTransaction").Return(tx)
	itemShopRepositoryMock.On("CommitTransaction", tx).Return(nil)
	itemShopRepositoryMock.On("RollbackTransaction", tx).Return(nil)

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(2), nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.SellingReq
		expected error
	}

	cases := []args{
		{
			label: "Selling item failed because the item is not enough",
			in: &_itemShopModel.SellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_itemShopException.ItemNotEnough{ItemID: 1},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Selling(c.in)
			assert.Error(t, err)
			assert.Equal(t, c.expected, err)
			assert.Nil(t, result)
		})
	}
}
