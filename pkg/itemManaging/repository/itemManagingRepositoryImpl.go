package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	itemManagingRepositoryImpl struct {
		db     *gorm.DB
		logger echo.Logger
	}
)

func NewItemManaginRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}
