package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/pkg/otp"
	"auth-service/internal/repository"
	"auth-service/pkg/otp"
	"auth-service/pkg/password"
	"auth-service/pkg/sms"
	"auth-service/pkg/token"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type AuthService struct {
	UserRepo    *repository.UserRepository
	ProfileRepo *repository.ProfileRepository
	OTPRepo     *repository.OTPRepository
	SMSSender   sms.Sender
	JWTSecret   string
}

func NewAuthService(
	userRepo *repository.UserRepository,
	profileRepo *repository.ProfileRepository,
	otpRepo *repository.OTPRepository,
	smsSender sms.Sender,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		UserRepo:    userRepo,
		ProfileRepo: profileRepo,
		OTPRepo:     otpRepo,
		SMSSender:   smsSender,
		JWTSecret:   jwtSecret,
	}
}

func (s *AuthService) SendOTP(phone string) error {
	code, err := otp.Generate()
	if err != nil {
		return err
	}
	expiresAt := time.Now().Add(2 * time.Minute)

	if err := s.OTPRepo.Save(phone, code, expiresAt); err != nil {
		return err
	}

	return s.SMSSender.Send(phone, fmt.Sprintf("Your OTP code is: %s", code))

}

func (s *AuthService) Register(phone, code, username, rawPassword string) (string, *domain.User, error) {
	existing, err := s.UserRepo.FindByPhone(phone)
	if err != nil && existing != nil {
		return "", nil, errors.New("phone number already registered")
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", nil, err
	}

	otpRow, err := s.OTPRepo.GetLatestActive(phone, code)
	if err != nil {
		if error.IS(err, sql.ErrNoRows) {
			return "", nil, errors.New("invalid OTP code")
		}
		return "", nil, err
	}
	if time.Now().After(otpRow.ExpiresAt) {
		return "", nil, errors.New("OTP code expired")
	}
	hashed, err := password.Hash(rawPassword)
	if err != nil {
		return "", nil, err
	}

	user := &domain.User{
		Phone:           phone,
		Username:        username,
		PasswordHash:    hashed,
		IsPhoneVerified: true,
	}

	if err := s.UserRepo.Create(user); err != nil {
		return "", nil, err
	}

	if err := s.ProfileRepo.CreateEmpty(user.ID); err != nil {
		return "", nil, err
	}

	if err := s.OTPRepo.MarkUsed(otpRow.ID); err != nil {
		return "", nil, err
	}

	jwtToken, err := token.GenerateJWT(user.ID, user.Phone, s.JWTSecret)
	if err != nil {
		return "", nil, err
	}

	return jwtToken, user, nil
}

func (s *AuthService) Login(phone, rawPassword string) (string, *domain.User, error) {
	user, err := s.UserRepo.FindByPhone(phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil, errors.New("invalid phone or password")
		}
		return "", nil, err
	}
	if err := password.Check(user.PasswordHash, rawPassword); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	jwtToken, err := token.GenerateJWT(user.ID, user.Phone, s.JWTSecret)
	if err != nil {
		return "", nil, err
	}

	return jwtToken, user, nil

}
