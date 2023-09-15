package repository

import (
	"context"
	"hostel-service/internal/authentication/domain"

	"gorm.io/gorm"
)

func NewAuthenticationAdapter(db *gorm.DB) *AuthenticationAdapter {
	return &AuthenticationAdapter{DB: db}
}

type AuthenticationAdapter struct {
	DB *gorm.DB
}

func (r *AuthenticationAdapter) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	r.DB.Where("username = ?", username).First(&user)
	// stmt, err := r.DB.Prepare(`select * from users where username = ? limit 1`)
	// if err != nil {
	// 	return nil, err
	// }
	// err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	// if err != nil {
	// 	if err != sql.ErrNoRows {
	// 		return nil, err
	// 	} else {
	// 		return nil, nil
	// 	}
	// }
	return &user, nil
}

func (r *AuthenticationAdapter) Create(ctx context.Context, user *domain.User) (int64, error) {
	// tx := q.GetTx(ctx)
	// stmt, err := tx.Prepare(`insert into users values(?,?,?,?,?)`)
	// if err != nil {
	// 	return 0, err
	// }
	// res, err := stmt.Exec(user.Id, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	// if err != nil {
	// 	return 0, err
	// }
	res := r.DB.Create(user)
	return res.RowsAffected, nil
}
