package repository

import (
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	"github.com/labstack/echo/v4"

	_adminException "github.com/guatom999/go-shop-api/pkg/admin/exception"
)

type (
	adminRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewAdminRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) AdminRepository {
	return &adminRepositoryImpl{db, logger}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.ConnectDatabase().Create(adminEntity).Scan(admin).Error; err != nil {
		r.logger.Errorf("Creating Admin Failed: %s", err.Error())
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.ConnectDatabase().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Find admin By ID Failed: %s", err.Error())
		return nil, &_adminException.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
