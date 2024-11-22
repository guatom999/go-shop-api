package service

import (
	_adminModel "github.com/guatom999/go-shop-api/pkg/admin/model"
	_playerModel "github.com/guatom999/go-shop-api/pkg/player/model"
)

type (
	OAuth2Service interface {
		PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
		AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
		IsUserArePlayer(playerID string) bool
		IsUserAreAdmin(adminID string) bool
	}
)
