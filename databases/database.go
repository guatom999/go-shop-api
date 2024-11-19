package databases

import "gorm.io/gorm"

type (
	Database interface {
		ConnectDatabase() *gorm.DB
	}
)
