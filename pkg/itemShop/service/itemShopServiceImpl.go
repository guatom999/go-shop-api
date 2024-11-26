package service

import (
	"github.com/guatom999/go-shop-api/entities"
	_inventoryRepository "github.com/guatom999/go-shop-api/pkg/inventory/repository"
	_itemShopException "github.com/guatom999/go-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/guatom999/go-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/guatom999/go-shop-api/pkg/itemShop/repository"
	_playerCoinModel "github.com/guatom999/go-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/guatom999/go-shop-api/pkg/playerCoin/repository"
	"github.com/labstack/echo/v4"
)

type (
	itemShopServiceImpl struct {
		itemShopRepository   _itemShopRepository.ItemShopRepository
		inventoryRepository  _inventoryRepository.InventoryRepository
		playerCoinRepository _playerCoinRepository.PlayerCoinRepository
		logger               echo.Logger
	}
)

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	playerCoinRepositoty _playerCoinRepository.PlayerCoinRepository,
	logger echo.Logger,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository:   itemShopRepository,
		inventoryRepository:  inventoryRepository,
		playerCoinRepository: playerCoinRepositoty,
		logger:               logger,
	}
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

func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {

	itemEntities, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntities.ToItemModel(), buyingReq.Quantity)

	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntities.Name,
		ItemDescription: itemEntities.Description,
		ItemPrice:       itemEntities.Price,
		ItemPicture:     itemEntities.Picture,
		Quantity:        buyingReq.Quantity,
	})
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Purchase history recorded: %s", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerId: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Infof("Player coin deducted: %d", playerCoin.Amount)

	inventoryEntity, err := s.inventoryRepository.Filling(
		tx,
		buyingReq.PlayerID,
		buyingReq.ItemID,
		int(buyingReq.Quantity),
	)
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Inventory Filled : %d", len(inventoryEntity))

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	return nil, nil
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

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Errorf("Player Coin is not enough %s", err.Error())
		return &_itemShopException.CoinNotEnough{}
	}

	return nil

}
