package service

import (
	"context"
	"encoding/json"
	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/port"
	rate_port "hostel-service/internal/rate/port"
	user_domain "hostel-service/internal/user/domain"
	user_port "hostel-service/internal/user/port"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

type PostService interface {
	GetPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	ESearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	SearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, postId string, userId string) (*domain.Post, error)
	GetCompare(ctx context.Context, post1Id string, post2Id string, userId string) ([]domain.Post, error)
	CheckCreatePost(ctx context.Context, userId string) (int64, error)
	CreatePost(ctx context.Context, post *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, post *domain.Post) (int64, error)
	UpdatePostStatus(ctx context.Context, userId string, status string) (int64, error)
	DeletePost(ctx context.Context, postId string) (int64, error)
}

func NewPostService(
	repository port.PostRepository,
	userRepo user_port.UserRepository,
	rateRepo rate_port.RateRepository,
	esClient *elasticsearch.Client,
) PostService {
	return &postService{
		repository: repository,
		userRepo:   userRepo,
		rateRepo:   rateRepo,
		esClient:   esClient,
	}
}

type postService struct {
	repository port.PostRepository
	userRepo   user_port.UserRepository
	rateRepo   rate_port.RateRepository
	esClient   *elasticsearch.Client
}

func (s *postService) GetPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	return s.repository.GetPosts(ctx, post, userId)
}

func (s *postService) SearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	return s.repository.GetPosts(ctx, post, userId)
}

func (s *postService) GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error) {
	if userId == "" {
		post := &domain.PostFilter{
			PageSize: 10,
			PageIdx:  0,
			Sort:     "view desc",
		}
		return s.repository.GetPosts(ctx, post, userId)
	}
	userSuggest, err := s.userRepo.GetUserSuggest(ctx, userId)
	if err != nil {
		return nil, 0, err
	}
	costFrom := userSuggest.Cost - 500000
	costTo := userSuggest.Cost + 500000
	capacityFrom := userSuggest.Capacity - 1
	capacityTo := userSuggest.Capacity + 1
	post := &domain.PostFilter{
		Province:     userSuggest.Province,
		District:     userSuggest.District,
		CostFrom:     costFrom,
		CostTo:       costTo,
		CapacityFrom: capacityFrom,
		CapacityTo:   capacityTo,
		PageSize:     10,
		PageIdx:      0,
		Sort:         "view desc",
	}
	posts, total, err := s.repository.GetPosts(ctx, post, userId)
	if total < 10 {
		post := &domain.PostFilter{
			Province: userSuggest.Province,
			PageSize: int(10 - total),
			PageIdx:  0,
			Sort:     "view asc",
		}
		addPosts, addTotal, _ := s.repository.GetPosts(ctx, post, userId)
		posts = append(posts, addPosts...)
		total += addTotal
	}
	return posts, total, err
}

func (s *postService) GetPostById(ctx context.Context, postId string, userId string) (*domain.Post, error) {
	res, err := s.repository.GetPostById(ctx, postId)
	if err != nil {
		return nil, err
	}
	res.RateInfo, err = s.rateRepo.GetPostRate(ctx, postId)
	if err != nil {
		return nil, err
	}
	if userId != "" {
		err = s.userRepo.UpdateUserSuggest(ctx, &user_domain.UpdateUserSuggest{
			Id:       userId,
			Province: res.Province,
			District: res.District,
			Cost:     res.Cost,
			Capacity: res.Capacity,
		})
	}
	return res, err
}

func (s *postService) GetCompare(ctx context.Context, post1Id string, post2Id string, userId string) ([]domain.Post, error) {
	res, err := s.repository.GetPostByIds(ctx, []string{post1Id, post2Id})
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *postService) CheckCreatePost(ctx context.Context, userId string) (int64, error) {
	user, err := s.userRepo.GetUserProfile(ctx, userId)
	if err != nil {
		return -1, err
	}

	if user.IsVerifiedPhone == nil || !*user.IsVerifiedPhone {
		return 0, nil
	}

	return 1, nil
}

func (s *postService) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
	post.Id = "post-" + uuid.NewString()
	post.CreatedAt = time.Now()
	post.EndedAt = time.Now().AddDate(0, 1, 0)
	post.Status = domain.PostActive
	return s.repository.CreatePost(ctx, post)
}

func (s *postService) UpdatePost(ctx context.Context, post *domain.Post) (int64, error) {
	t := time.Now()
	post.UpdatedAt = &t
	return s.repository.UpdatePost(ctx, post)
}

func (s *postService) UpdatePostStatus(ctx context.Context, userId string, status string) (int64, error) {
	return s.repository.UpdatePostStatus(ctx, userId, status)
}

func (s *postService) DeletePost(ctx context.Context, postId string) (int64, error) {
	post := &domain.Post{Id: postId}
	return s.repository.DeletePost(ctx, post)
}
func (s *postService) ESearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	posts, total, err := s.eSearchPosts(ctx, post)
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil

}

func (s *postService) eSearchPosts(ctx context.Context, post *domain.PostFilter) ([]domain.Post, int64, error) {
	elasticQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"name": post.Name,
						},
					},
					{
						"match": map[string]interface{}{
							"district": map[string]interface{}{
								"query": post.District,
								"boost": 2,
							},
						},
					},
					{
						"match": map[string]interface{}{
							"ward": post.Ward,
						},
					},
					{
						"range": map[string]interface{}{
							"cost": domain.Range{
								GTE:   post.CostFrom,
								LTE:   post.CostTo,
								Boost: 2,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"deposit": domain.Range{
								GTE: post.DepositFrom,
								LTE: post.DepositTo,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"capacity": domain.Range{
								GTE: post.Capacity - 1,
								LTE: post.Capacity + 1,
							},
						},
					},
				},
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"status": "A",
						},
					},
					{
						"match": map[string]interface{}{
							"province": map[string]interface{}{
								"query": strings.Replace(post.Province, "Thành phố ", "", -1),
								"boost": 3,
							},
						},
					},
				},
			},
		},
		// "filter": map[string]interface{}{
		// 	"range": map[string]interface{}{
		// 		"_score": domain.Range{
		// 			GTE: 1,
		// 		},
		// 	},
		// },
		// "sort": []map[string]interface{}{
		// 	{
		// 		"createdAt": map[string]interface{}{
		// 			"order": "asc", // Hoặc "desc" nếu bạn muốn sắp xếp giảm dần.
		// 		},
		// 	},
		// },
	}

	if len(post.Utilities) > 0 {
		elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"terms_set": map[string]interface{}{
					"utilities": map[string]interface{}{
						"terms":                post.Utilities,
						"minimum_should_match": len(post.Utilities) - 1,
					},
				},
			},
		)
	}
	elasticQuery["size"] = 100
	elasticQueryJSON, err := json.Marshal(elasticQuery)
	if err != nil {
		return nil, 0, err
	}
	// Gửi yêu cầu tìm kiếm đến Elasticsearch
	req := esapi.SearchRequest{
		Index: []string{"post_index"},
		Body:  strings.NewReader(string(elasticQueryJSON)),
	}

	res, err := req.Do(context.Background(), s.esClient)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, err
	}

	var result domain.SearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, err
	}

	posts := []domain.Post{}
	for _, post := range result.Hits.Hits {
		posts = append(posts, post.Source)
	}

	return posts, int64(len(posts)), err
}
