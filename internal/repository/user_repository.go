package repository

import (
	"auth-service/internal/domain"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `
	INSERT INTO users(phone, username, password_hash, is_phone_verified)
	VALUES($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.DB.QueryRow(
		query,
		user.Phone,
		user.Username,
		user.PasswordHash,
		user.IsPhoneVerified,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) FindByPhone(phone string) (*domain.User, error) {
	query := `
	SELECT id, phone, username, password_hash, is_phone_verified, created_at
	FROM users
	WHERE phone = $1
	`

	u := &domain.User{}
	err := r.DB.QueryRow(query, phone).Scan(
		&u.ID,
		&u.Phone,
		&u.Username,
		&u.PasswordHash,
		&u.IsPhoneVerified,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	query := `
	SELECT id, phone, username, password_hash, is_phone_verified, created_at
	FROM users
	WHERE id = $1
	`

	u := &domain.User{}
	err := r.DB.QueryRow(query, id).Scan(
		&u.ID,
		&u.Phone,
		&u.Username,
		&u.PasswordHash,
		&u.IsPhoneVerified,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
