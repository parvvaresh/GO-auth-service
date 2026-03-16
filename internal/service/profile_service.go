package service

import "auth-service/internal/domain"

type ProfileRepo interface {
	GetByUserID(userID int) (*domain.Profile, error)
	Update(profile *domain.Profile) error
}

type ProfileService struct {
	Repo ProfileRepo
}

func NewProfileService(repo ProfileRepo) *ProfileService {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) Get(userID int) (*domain.Profile, error) {
	return s.Repo.GetByUserID(userID)
}

func (s *ProfileService) Update(profile *domain.Profile) error {
	return s.Repo.Update(profile)
}
