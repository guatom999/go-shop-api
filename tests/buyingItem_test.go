package tests

import (
	"testing"

	"github.com/guatom999/go-shop-api/entities"
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"

	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"

	_itemShopService "github.com/guatom999/go-shop-api/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

func TestItemBuyingSuccess(t *testing.T) {
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
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P0001").Return(&_playerCoinModel.PlayerCoinShowing{
		PlayerID: "P0001",
		Coin:     10000,
	}, nil)

	itemShopRepositoryMock.On("PurchaseHistoryRecording", &entities.PurchaseHistory{
		PlayerID:        "P0001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        4,
		IsBuying:        true,
	}, tx).Return(&entities.PurchaseHistory{
		PlayerID:        "P0001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        4,
		IsBuying:        true,
	}, nil)

	playerCoinRepositoryMock.On("CoinAdding", tx, &entities.PlayerCoin{
		PlayerId: "P0001",
		Amount:   -4000,
	}).Return(&entities.PlayerCoin{
		PlayerId: "P0001",
		Amount:   -4000,
	}, nil)

	inventoryRepositoryMock.On("Filling", tx, "P0001", uint64(1), int(4)).Return(
		[]*entities.Inventory{
			{
				PlayerID: "P0001",
				ItemID:   1,
			},
			{
				PlayerID: "P0001",
				ItemID:   1,
			},
			{
				PlayerID: "P0001",
				ItemID:   1,
			},
			{
				PlayerID: "P0001",
				ItemID:   1,
			},
		},
		nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected *_playerCoinModel.PlayerCoin
	}

	cases := []args{
		{
			label: "Success buying item",
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P0001",
				ItemID:   1,
				Quantity: 4,
			},
			expected: &_playerCoinModel.PlayerCoin{
				PlayerID: "P0001",
				Amount:   -4000,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)
		})
	}

}

func TestItemBuyingFail(t *testing.T) {
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

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&_playerCoinModel.PlayerCoinShowing{
		PlayerID: "P001",
		Coin:     2000,
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected error
	}

	cases := []args{
		{
			"Test Find Item Failed",
			&_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_itemShopException.CoinNotEnough{},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.Nil(t, result)
			assert.Error(t, err)
			assert.EqualValues(t, c.expected, err)
		})
	}
}
