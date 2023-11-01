package repository

import (
	"context"
	"hostel-service/internal/package/util"
	"hostel-service/internal/rate/domain"

	"gorm.io/gorm"
)

func NewRateAdapter(db *gorm.DB) *RateAdapter {
	return &RateAdapter{DB: db}
}

type RateAdapter struct {
	DB *gorm.DB
}

func (r *RateAdapter) GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	postRate := &domain.PostRateInfo{}
	res := r.DB.Table("post_rate_info").Where("post_id = ?", postId).Find(postRate)
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.DB.Table("rates").Select("rates.*, users.display_name, users.avatar_url").
		Joins("left join users on users.id = rates.user_id").
		Where("post_id = ?", postId).Scan(&postRate.RateList)
	postRate.AvgRate = util.RoundFloat(float64(*postRate.Star1+2*(*postRate.Star2)+3*(*postRate.Star3)+4*(*postRate.Star4)+5*(*postRate.Star5))/float64(*postRate.Total), 1)
	return postRate, res.Error
}

func (r *RateAdapter) GetSimplePostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	postRate := &domain.PostRateInfo{}
	res := r.DB.Table("post_rate_info").Where("post_id = ?", postId).Find(postRate)
	return postRate, res.Error
}

func (r *RateAdapter) UpdatePostRateInfo(ctx context.Context, rateInfo *domain.PostRateInfo) error {
	res := r.DB.Table("post_rate_info").Where("post_id = ?", rateInfo.PostId).Model(&rateInfo).Updates(rateInfo)
	return res.Error
}

func (r *RateAdapter) GetRate(ctx context.Context, postId string, userId string) (*domain.Rate, error) {
	rate := &domain.Rate{}
	res := r.DB.Table("rates").Select("rates.*").Where("post_id = ? and user_id = ?", postId, userId).Scan(rate)
	return rate, res.Error
}

func (r *RateAdapter) CreateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	res := r.DB.Table("rates").Create(rate)
	return res.RowsAffected, res.Error
}

func (r *RateAdapter) UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	res := r.DB.Table("rates").Model(&rate).Updates(rate)
	return res.RowsAffected, res.Error
}
