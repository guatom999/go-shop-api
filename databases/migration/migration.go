package main

import (
	"fmt"

	"github.com/guatom999/go-shop-api/config"
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectDatabase())

	tx := db.ConnectDatabase().Begin()

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseMigration(tx)

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
