package main

import (
	"github.com/guatom999/go-shop-api/config"
	"github.com/guatom999/go-shop-api/databases"
	"github.com/guatom999/go-shop-api/server"
)

func main() {

	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
