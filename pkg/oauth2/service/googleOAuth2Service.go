package service

import (
	"github.com/guatom999/go-shop-api/entities"
	_adminModel "github.com/guatom999/go-shop-api/pkg/admin/model"
	_adminRepository "github.com/guatom999/go-shop-api/pkg/admin/repository"
	_playerModel "github.com/guatom999/go-shop-api/pkg/player/model"
	_playerRepository "github.com/guatom999/go-shop-api/pkg/player/repository"
)

type (
	googleOAuth2Service struct {
		playerRepository _playerRepository.PlayerRepository
		adminRepository  _adminRepository.AdminRepository
	}
)

func NewGoogleOAuthService(playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {

	if !s.IsUserArePlayer(playerCreatingReq.ID) {

		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Email:  playerCreatingReq.Email,
			Name:   playerCreatingReq.Name,
			Avatar: playerCreatingReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {
	if !s.IsUserAreAdmin(adminCreatingReq.ID) {
		adminEntity := &entities.Admin{
			ID:     adminCreatingReq.ID,
			Email:  adminCreatingReq.Email,
			Name:   adminCreatingReq.Name,
			Avatar: adminCreatingReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}

	}

	return nil
}

func (s *googleOAuth2Service) IsUserArePlayer(playerID string) bool {

	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsUserAreAdmin(adminID string) bool {

	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}

	return admin != nil
}
