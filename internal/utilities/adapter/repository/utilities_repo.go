package repository

import (
	"context"
	"hostel-service/internal/utilities/domain"

	"gorm.io/gorm"
)

func NewUtilitiesAdapter(db *gorm.DB) *UtilitiesAdapter {
	return &UtilitiesAdapter{DB: db}
}

type UtilitiesAdapter struct {
	DB *gorm.DB
}

func (r *UtilitiesAdapter) GetAllUtilities(ctx context.Context) ([]domain.GetUtilities, error) {
	var utilities []domain.GetUtilities
	r.DB.Table("utilities").Find(&utilities)
	return utilities, nil
}

func (r *UtilitiesAdapter) CreateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Create(utilities)
	return res.RowsAffected, res.Error
}

func (r *UtilitiesAdapter) UpdateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Model(&utilities).Updates(utilities)
	return res.RowsAffected, res.Error
}

func (r *UtilitiesAdapter) DeleteUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	res := r.DB.Table("utilities").Delete(utilities)
	return res.RowsAffected, res.Error
}
