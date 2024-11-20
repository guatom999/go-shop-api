package repository

import "github.com/guatom999/go-shop-api/entities"

type (
	ItemManagingRepository interface {
		Creating(itemEntity *entities.Item) (*entities.Item, error)
	}
)
