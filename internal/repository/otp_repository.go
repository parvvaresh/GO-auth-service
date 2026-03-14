package repository

import (
	"auth-service/internal/domain"
	"database/sql"
	"time"
)

type OTPRepository struct {
	DB *sql.DB
}

func (r *OTPRepository) Save(phone, code string, expiresAt time.Time) error {
	query := `
	INSERT INTO otp_codes(phone, code, expires_at, used)
	VALUES($1, $2, $3, FALSE)
	`
	_, err := r.DB.Exec(query, phone, code, expiresAt)
	return err
}

func (r *OTPRepository) GetLatestActive(phone, code string) (*domain.OTPCode, error) {
	query := `
	SELECT id, phone, code, expires_at, used, created_at
	FROM otp_codes
	WHERE phone = $1 AND code = $2 AND used = FALSE
	ORDER BY id DESC
	LIMIT 1
	`

	o := &domain.OTPCode{}
	err := r.DB.QueryRow(query, phone, code).Scan(
		&o.ID,
		&o.Phone,
		&o.Code,
		&o.ExpiresAt,
		&o.Used,
		&o.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *OTPRepository) MarkUsed(id int) error {
	_, err := r.DB.Exec(`UPDATE otp_codes SET used = TRUE WHERE id = $1`, id)
	return err
}
