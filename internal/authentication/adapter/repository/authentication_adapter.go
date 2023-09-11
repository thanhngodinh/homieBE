package repository

import (
	"context"
	"database/sql"
	"hostel-service/internal/authentication/domain"

	q "github.com/core-go/sql"
)

func NewAuthenticationAdapter(db *sql.DB) *AuthenticationAdapter {
	return &AuthenticationAdapter{DB: db}
}

type AuthenticationAdapter struct {
	DB *sql.DB
}

func (r *AuthenticationAdapter) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	stmt, err := r.DB.Prepare(`select * from users where username = ? limit 1`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return &user, nil
}

func (r *AuthenticationAdapter) Create(ctx context.Context, user *domain.User) (int64, error) {
	tx := q.GetTx(ctx)
	stmt, err := tx.Prepare(`insert into users values(?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.Id, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
