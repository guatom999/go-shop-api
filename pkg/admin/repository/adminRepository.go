package repository

import "github.com/guatom999/go-shop-api/entities"

type (
	AdminRepository interface {
		Creating(adminEntity *entities.Admin) (*entities.Admin, error)
		FindByID(adminID string) (*entities.Admin, error)
	}
)
