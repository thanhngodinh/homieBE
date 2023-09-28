package repository

import (
	"context"
	"hostel-service/internal/user/domain"

	"gorm.io/gorm"
)

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{DB: db}
}

type UserAdapter struct {
	DB *gorm.DB
}

func (r *UserAdapter) UpdateUserSuggest(ctx context.Context, us *domain.UpdateUserSuggest) error {
	res := r.DB.Table("users").Updates(us)
	return res.Error
}

func (r *UserAdapter) GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error) {
	res := &domain.UserSuggest{}
	r.DB.Table("users").Where("id = ?", userId).First(res)
	return res, nil
}

func (r *UserAdapter) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	r.DB.Where("username = ?", username).First(&user)
	return &user, nil
}

func (r *UserAdapter) Create(ctx context.Context, user *domain.User) (int64, error) {
	res := r.DB.Create(user)
	return res.RowsAffected, nil
}
