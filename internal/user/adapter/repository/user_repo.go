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
	user := &domain.User{}
	r.DB.Where("username = ?", username).First(user)
	return user, nil
}

func (r *UserAdapter) GetById(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}
	r.DB.Where("id = ?", userId).First(user)
	return user, nil
}

func (r *UserAdapter) Create(ctx context.Context, user *domain.User) error {
	res := r.DB.Create(user)
	return res.Error
}

func (r *UserAdapter) UpdatePassword(ctx context.Context, userId string, newPassword string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"password": newPassword})
	return res.Error
}
