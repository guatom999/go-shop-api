package repository

import (
	entities "github.com/guatom999/go-shop-api/entities"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ItemShopRepositoryMock struct {
	mock.Mock
}

func (m *ItemShopRepositoryMock) TransactionBegin() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *ItemShopRepositoryMock) TransactionRollback(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *ItemShopRepositoryMock) TransactionCommit(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *ItemShopRepositoryMock) FindByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	args := m.Called(itemFilter)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	args := m.Called(itemFilter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *ItemShopRepositoryMock) PurchaseHistoryRecording(tx *gorm.DB, purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	args := m.Called(purchasingEntity, tx)
	return args.Get(0).(*entities.PurchaseHistory), args.Error(1)
}