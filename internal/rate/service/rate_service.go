package service

import (
	"context"
	"encoding/json"
	"hostel-service/internal/rate/domain"
	"hostel-service/internal/rate/port"
	"hostel-service/pkg/util"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type RateService interface {
	GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error)
	CreateRate(ctx context.Context, rate *domain.Rate) (int64, error)
	UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error)
}

func NewRateService(
	repository port.RateRepository,
	esClient *elasticsearch.Client,

) RateService {
	return &rateService{
		repository: repository,
		esClient:   esClient,
	}
}

type rateService struct {
	repository port.RateRepository
	esClient   *elasticsearch.Client
}

func (s *rateService) GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	return s.repository.GetPostRate(ctx, postId)
}

func (s *rateService) CreateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	rate.CreatedAt = time.Now()
	res, err := s.repository.CreateRate(ctx, rate)
	if err != nil {
		return 0, err
	}

	rateInfo, err := s.repository.GetSimplePostRate(ctx, rate.PostId)
	if err != nil {
		return 0, err
	}
	*rateInfo.Total++
	switch rate.Star {
	case 1:
		*rateInfo.Star1++
	case 2:
		*rateInfo.Star2++
	case 3:
		*rateInfo.Star3++
	case 4:
		*rateInfo.Star4++
	case 5:
		*rateInfo.Star5++
	}
	rateInfo.AvgRate = util.RoundFloat(float64(*rateInfo.Star1+2*(*rateInfo.Star2)+3*(*rateInfo.Star3)+4*(*rateInfo.Star4)+5*(*rateInfo.Star5))/float64(*rateInfo.Total), 1)

	err = s.updateElasticRate(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	err = s.repository.UpdatePostRateInfo(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *rateService) UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	oldRate, err := s.repository.GetRate(ctx, rate.PostId, rate.UserId)
	if err != nil {
		return 0, err
	}

	t := time.Now()
	rate.CreatedAt = oldRate.CreatedAt
	rate.UpdatedAt = &t
	res, err := s.repository.UpdateRate(ctx, rate)
	if err != nil {
		return 0, err
	}

	rateInfo, err := s.repository.GetSimplePostRate(ctx, rate.PostId)
	if err != nil {
		return 0, err
	}

	switch oldRate.Star {
	case 1:
		*rateInfo.Star1--
	case 2:
		*rateInfo.Star2--
	case 3:
		*rateInfo.Star3--
	case 4:
		*rateInfo.Star4--
	case 5:
		*rateInfo.Star5--
	}

	switch rate.Star {
	case 1:
		*rateInfo.Star1++
	case 2:
		*rateInfo.Star2++
	case 3:
		*rateInfo.Star3++
	case 4:
		*rateInfo.Star4++
	case 5:
		*rateInfo.Star5++
	}
	rateInfo.AvgRate = util.RoundFloat(float64(*rateInfo.Star1+2*(*rateInfo.Star2)+3*(*rateInfo.Star3)+4*(*rateInfo.Star4)+5*(*rateInfo.Star5))/float64(*rateInfo.Total), 1)

	err = s.updateElasticRate(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	err = s.repository.UpdatePostRateInfo(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *rateService) updateElasticRate(ctx context.Context, rateInfo *domain.PostRateInfo) error {
	updateRequest := map[string]interface{}{
		"doc": map[string]interface{}{
			"rateInfo": rateInfo,
		},
	}
	elasticJSON, err := json.Marshal(updateRequest)
	if err != nil {
		return err
	}
	// Thực hiện cập nhật trong Elasticsearch
	req := esapi.UpdateRequest{
		Index:      "post_index",
		DocumentID: rateInfo.PostId,
		Body:       strings.NewReader(string(elasticJSON)),
	}

	eRes, err := req.Do(ctx, s.esClient)
	if err != nil {
		return err
	}
	defer eRes.Body.Close()

	return nil
}
