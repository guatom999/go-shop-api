package service

import (
	_adminRepository "github.com/guatom999/go-shop-api/pkg/admin/repository"
	_playerRepository "github.com/guatom999/go-shop-api/pkg/player/repository"
)

type (
	googleOAuth2Service struct {
		playerRepository _playerRepository.PlayerRepository
		adminRepository  _adminRepository.AdminRepository
	}
)

func NewGoogleOAuthRepository(playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}
