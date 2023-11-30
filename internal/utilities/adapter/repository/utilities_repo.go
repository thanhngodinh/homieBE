package repository

import (
	"context"
	"hostel-service/internal/utilities/domain"

	"gorm.io/gorm"
)

func NewUtilitiesRepo(db *gorm.DB) *UtilitiesRepo {
	return &UtilitiesRepo{DB: db}
}

type UtilitiesRepo struct {
	DB *gorm.DB
}

func (r *UtilitiesRepo) GetAllUtilities(ctx context.Context) ([]domain.GetUtilities, error) {
	var utilities []domain.GetUtilities
	r.DB.Table("utilities").Find(&utilities)
	return utilities, nil
}

func (r *UtilitiesRepo) CreateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Create(utilities)
	return res.RowsAffected, res.Error
}

func (r *UtilitiesRepo) UpdateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Model(&utilities).Updates(utilities)
	return res.RowsAffected, res.Error
}

func (r *UtilitiesRepo) DeleteUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Delete(utilities)
	return res.RowsAffected, res.Error
}
