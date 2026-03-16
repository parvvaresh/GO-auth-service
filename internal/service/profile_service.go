package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type ProfileService struct {
	ProfileRepo *repository.ProfileRepository
}

func NewProfileService(profileRepo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{ProfileRepo: profileRepo}
}

func (s *ProfileService) Get(userID int) (*domain.Profile, error) {
	return s.ProfileRepo.GetByUserID(userID)
}

func (s *ProfileService) Update(profile *domain.Profile) error {
	return s.ProfileRepo.Update(profile)
}
