package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	_playerException "github.com/guatom999/go-shop-api/pkg/player/exception"
	"github.com/labstack/echo/v4"
)

type (
	playerRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewPlayerRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) PlayerRepository {
	return &playerRepositoryImpl{db, logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {

	player := new(entities.Player)

	if err := r.db.ConnectDatabase().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Errorf("Creating Player Failed: %s", err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {

	player := new(entities.Player)

	if err := r.db.ConnectDatabase().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Find Player By ID Failed: %s", err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
