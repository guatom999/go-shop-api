package service

import (
	"github.com/guatom999/go-shop-api/entities"
	_inventoryModel "github.com/guatom999/go-shop-api/pkg/inventory/model"
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
)

type (
	inventoryServiceImpl struct {
		inventoryRepository _inventoryRepository.InventoryRepository
		itemShopRepository  _itemShopRepository.ItemShopRepository
	}
)

func NewInventoryServiceImpl(inventoryRepository _inventoryRepository.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{inventoryRepository, itemShopRepository}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {

	inventoryEntities, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventoryEntities)

	return s.buildInventoryListingResult(uniqueItemWithQuantityCounterList), nil
}

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(inventoryEntities []*entities.Inventory) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList

}

func (s *inventoryServiceImpl) buildInventoryListingResult(uniqueItemWithQuantityConterList []_inventoryModel.ItemQuantityCounting) []*_inventoryModel.Inventory {

	uniqueItemIDList := s.getItemID(uniqueItemWithQuantityConterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityConterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}

	return results
}

func (s *inventoryServiceImpl) getItemID(uniqueItemWithQuantityConterList []_inventoryModel.ItemQuantityCounting) []uint64 {

	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityConterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(uniqueItemWithQuantityConterList []_inventoryModel.ItemQuantityCounting) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantityConterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity

}
