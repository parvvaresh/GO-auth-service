package repository

import (
	"auth-service/internal/domain"
	"database/sql"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) CreateEmpty(userID int) error {
	query := `
	INSERT INTO profiles(user_id, first_name, last_name, bio)
	VALUES($1, '', '', '')
	ON CONFLICT (user_id) DO NOTHING
	`
	_, err := r.DB.Exec(query, userID)
	return err
}

func (r *ProfileRepository) GetByUserID(userID int) (*domain.Profile, error) {
	query := `
	SELECT id, user_id, first_name, last_name, bio
	FROM profiles
	WHERE user_id = $1
	`

	p := &domain.Profile{}
	err := r.DB.QueryRow(query, userID).Scan(
		&p.ID,
		&p.UserID,
		&p.FirstName,
		&p.LastName,
		&p.Bio,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProfileRepository) Update(profile *domain.Profile) error {
	query := `
	UPDATE profiles
	SET first_name = $1, last_name = $2, bio = $3
	WHERE user_id = $4
	`
	_, err := r.DB.Exec(query,
		profile.FirstName,
		profile.LastName,
		profile.Bio,
		profile.UserID,
	)
	return err
}
