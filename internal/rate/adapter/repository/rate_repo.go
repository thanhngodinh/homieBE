package repository

import (
	"context"
	"hostel-service/internal/rate/domain"

	"gorm.io/gorm"
)

func NewRateRepo(db *gorm.DB) *RateRepo {
	return &RateRepo{DB: db}
}

type RateRepo struct {
	DB *gorm.DB
}

func (r *RateRepo) GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	postRate := &domain.PostRateInfo{}
	res := r.DB.Table("post_rate_info").Where("post_id = ?", postId).Find(postRate)
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.DB.Table("rates").Select("rates.*, users.display_name, users.avatar_url").
		Joins("left join users on users.id = rates.user_id").
		Where("post_id = ?", postId).Scan(&postRate.RateList)
	return postRate, res.Error
}

func (r *RateRepo) GetSimplePostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	postRate := &domain.PostRateInfo{}
	res := r.DB.Table("post_rate_info").Where("post_id = ?", postId).Find(postRate)

	return postRate, res.Error
}

func (r *RateRepo) UpdatePostRateInfo(ctx context.Context, rateInfo *domain.PostRateInfo) error {
	res := r.DB.Table("post_rate_info").Where("post_id = ?", rateInfo.PostId).Model(&rateInfo).Updates(rateInfo)
	return res.Error
}

func (r *RateRepo) GetRate(ctx context.Context, postId string, userId string) (*domain.Rate, error) {
	rate := &domain.Rate{}
	res := r.DB.Table("rates").Select("rates.*").Where("post_id = ? and user_id = ?", postId, userId).Scan(rate)
	return rate, res.Error
}

func (r *RateRepo) CreateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	res := r.DB.Table("rates").Create(rate)
	return res.RowsAffected, res.Error
}

func (r *RateRepo) UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	res := r.DB.Table("rates").Model(&rate).Updates(rate)
	return res.RowsAffected, res.Error
}
