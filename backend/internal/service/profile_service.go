package service

import (
	"context"

	"github.com/Raffayet/data-merging/backend/internal/domain"
	"github.com/Raffayet/data-merging/backend/internal/repository"
)

// ProfileService defines methods for working with profiles
type ProfileService interface {
	GetProfiles(ctx context.Context) ([]domain.Profile, error)
}

// ProfileServiceImpl implements the ProfileService interface
type ProfileServiceImpl struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) *ProfileServiceImpl {
	return &ProfileServiceImpl{repo: repo}
}

// GetProfiles fetches profiles from the repository
func (s *ProfileServiceImpl) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	return s.repo.FetchProfiles(ctx)
}
