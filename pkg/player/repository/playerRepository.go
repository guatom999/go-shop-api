package repository

import "github.com/guatom999/go-shop-api/entities"

type (
	PlayerRepository interface {
		Creating(playerEntity *entities.Player) (*entities.Player, error)
		FindByID(playerID string) (*entities.Player, error)
	}
)
