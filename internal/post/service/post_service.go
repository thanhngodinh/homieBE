package service

import (
	"context"
	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/port"
	rate_port "hostel-service/internal/rate/port"
	user_domain "hostel-service/internal/user/domain"
	user_port "hostel-service/internal/user/port"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
)

type PostService interface {
	GetPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	SearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, postId string, userId string) (*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, post *domain.Post) (int64, error)
	DeletePost(ctx context.Context, postId string) (int64, error)
}

func NewPostService(
	repository port.PostRepository,
	userRepo user_port.UserRepository,
	rateRepo rate_port.RateRepository,
) PostService {
	return &postService{
		repository: repository,
		userRepo:   userRepo,
		rateRepo:   rateRepo,
	}
}

type postService struct {
	repository          port.PostRepository
	userRepo            user_port.UserRepository
	rateRepo            rate_port.RateRepository
	ElasticsearchClient *elasticsearch.Client
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
		Province:     &userSuggest.Province,
		District:     &userSuggest.District,
		CostFrom:     &costFrom,
		CostTo:       &costTo,
		CapacityFrom: &capacityFrom,
		CapacityTo:   &capacityTo,
		PageSize:     10,
		PageIdx:      0,
		Sort:         "view desc",
	}
	posts, total, err := s.repository.GetPosts(ctx, post, userId)
	if total < 10 {
		post := &domain.PostFilter{
			Province: &userSuggest.Province,
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

func (s *postService) DeletePost(ctx context.Context, postId string) (int64, error) {
	post := &domain.Post{Id: postId}
	return s.repository.DeletePost(ctx, post)
}

// func (s *postService) IndexToElasticsearch(record domain.Post) error {
// 	doc := map[string]interface{}{
// 		"id":          record.ID,
// 		"name":        record.Name,
// 		"description": record.Description,
// 		// Thêm các trường khác
// 	}

// 	// Lập chỉ mục dữ liệu vào Elasticsearch
// 	docJSON := fmt.Sprintf(`{
// 		"index": {
// 			"_index": "your_index",
// 			"_id": "%d"
// 		}
// 	}`, record.ID)

// 	req := esapi.IndexRequest{
// 		Index:      "your_index",
// 		DocumentID: fmt.Sprintf("%d", record.ID),
// 		Body:       strings.NewReader(docJSON),
// 	}
// 	reqop, err := req.Do(context.Background(), s.ElasticsearchClient)
// 	if err != nil {
// 		return err
// 	}
// 	if reqop.Error != nil {
// 		return reqop.Error
// 	}
// 	return nil
// }
