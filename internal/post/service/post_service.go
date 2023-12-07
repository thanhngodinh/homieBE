package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/port"
	rate_port "hostel-service/internal/rate/port"
	user_domain "hostel-service/internal/user/domain"
	user_port "hostel-service/internal/user/port"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

type PostService interface {
	ESearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	SearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, postId string, userId string) (*domain.Post, error)
	GetCompare(ctx context.Context, post1Id string, post2Id string, userId string) ([]domain.Post, error)
	CheckCreatePost(ctx context.Context, userId string) (int64, error)
	CreatePost(ctx context.Context, post *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, post *domain.UpdatePostReq) (int64, error)
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

func (s *postService) SearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	return s.repository.GetPosts(ctx, post, userId)
}

func (s *postService) GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error) {
	if userId == "" {
		post := &domain.PostFilter{
			PageSize: 5,
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
		PageSize:     5,
		PageIdx:      0,
		Sort:         "view desc",
	}
	posts, total, err := s.eSearchPosts(ctx, post)
	// if total < 10 {
	// 	post := &domain.PostFilter{
	// 		Province: userSuggest.Province,
	// 		PageSize: int(10 - total),
	// 		PageIdx:  0,
	// 		Sort:     "view asc",
	// 	}
	// 	addPosts, addTotal, _ := s.repository.GetPosts(ctx, post, userId)
	// 	posts = append(posts, addPosts...)
	// 	total += addTotal
	// }
	return posts, total, err
}

func (s *postService) GetPostById(ctx context.Context, postId string, userId string) (*domain.Post, error) {
	res, err := s.repository.GetPostById(ctx, postId, userId)
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

	err := s.indexPost(s.esClient, post)
	if err != nil {
		return -1, err
	}

	return s.repository.CreatePost(ctx, post)
}

func (s *postService) UpdatePost(ctx context.Context, post *domain.UpdatePostReq) (int64, error) {
	t := time.Now()
	post.UpdatedAt = &t

	// Thực hiện cập nhật trong Elasticsearch
	err := s.updateElasticPost(ctx, s.esClient, post)
	if err != nil {
		return -1, err
	}

	return s.repository.UpdatePost(ctx, post)
}

func (s *postService) DeletePost(ctx context.Context, postId string) (int64, error) {
	post := &domain.Post{Id: postId}

	deleteResponse, err := s.esClient.Delete("post_index", postId)
	if err != nil || deleteResponse.IsError() {
		return -1, fmt.Errorf("Error deleting post from Elasticsearch: %v | %v", err, deleteResponse.String())
	}

	// deleteAllRecords(s.esClient, "post_index")

	return s.repository.DeletePost(ctx, post)
}
func (s *postService) ESearchPosts(ctx context.Context, post *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	posts, total, err := s.eSearchPosts(ctx, post)
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (s *postService) eSearchPosts(ctx context.Context, filter *domain.PostFilter) ([]domain.Post, int64, error) {
	currentDate := time.Now()
	elasticQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"name": filter.Name,
						},
					},
					{
						"match": map[string]interface{}{
							"district": map[string]interface{}{
								"query": filter.District,
								"boost": 2,
							},
						},
					},
					{
						"match": map[string]interface{}{
							"ward": filter.Ward,
						},
					},
					{
						"range": map[string]interface{}{
							"cost": domain.Range{
								GTE:   filter.CostFrom,
								LTE:   filter.CostTo,
								Boost: 2,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"deposit": domain.Range{
								GTE: filter.DepositFrom,
								LTE: filter.DepositTo,
							},
						},
					},
					{
						"range": map[string]interface{}{
							"capacity": domain.Range{
								GTE: filter.Capacity - 1,
								LTE: filter.Capacity + 1,
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
						"range": map[string]interface{}{
							"endedAt": map[string]interface{}{
								"gte": currentDate.Format(time.RFC3339),
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
	if len(filter.Province) > 0 {
		elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"match": map[string]interface{}{
					"province": map[string]interface{}{
						"query": strings.Replace(filter.Province, "Thành phố ", "", -1),
						"boost": 3,
					},
				},
			},
		)
	}
	if len(filter.Utilities) > 0 {
		elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			elasticQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"terms_set": map[string]interface{}{
					"utilities": map[string]interface{}{
						"terms":                filter.Utilities,
						"minimum_should_match": 1,
					},
				},
			},
		)
	}
	if filter.PageSize >= 100 || filter.PageSize <= 0 {
		filter.PageSize = 5
	}
	if filter.PageIdx <= 0 {
		filter.PageIdx = 1
	}
	elasticQuery["size"] = filter.PageSize
	elasticQuery["from"] = (filter.PageIdx - 1) * filter.PageSize
	elasticQueryJSON, err := json.Marshal(elasticQuery)
	if err != nil {
		return nil, 0, err
	}

	fmt.Println(string(elasticQueryJSON))
	// Gửi yêu cầu tìm kiếm đến Elasticsearch
	req := esapi.SearchRequest{
		Index: []string{"post_index"},
		Body:  strings.NewReader(string(elasticQueryJSON)),
	}

	res, err := req.Do(ctx, s.esClient)
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
	// fmt.Println(result)
	posts := []domain.Post{}
	for _, post := range result.Hits.Hits {
		posts = append(posts, post.Source)
	}

	return posts, result.Hits.Total.Value, err
}

func (s *postService) indexPost(es *elasticsearch.Client, post *domain.Post) error {
	docID := post.Id
	indexName := "post_index"

	// Chuyển đổi struct thành JSON
	doc, err := json.Marshal(post)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return err
	}

	// Gửi dữ liệu vào Elasticsearch
	res, err := es.Index(indexName, bytes.NewReader(doc), es.Index.WithDocumentID(docID))
	if err != nil {
		log.Printf("Error indexing document: %v", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document: %s", res.String())
		return err
	}

	log.Printf("Indexed document with ID %s to index %s", docID, indexName)
	return nil
}

func (s *postService) updateElasticPost(ctx context.Context, es *elasticsearch.Client, post *domain.UpdatePostReq) error {
	elasticJSON, err := json.Marshal(post)
	if err != nil {
		return err
	}
	req := esapi.UpdateRequest{
		Index:      "post_index",
		DocumentID: post.Id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, elasticJSON))),
	}

	eRes, err := req.Do(ctx, s.esClient)
	if err != nil {
		return err
	}
	defer eRes.Body.Close()

	return nil
}

func (s *postService) deleteAllRecords(es *elasticsearch.Client, indexName string) error {
	// Tạo một truy vấn match_all
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	doc, err := json.Marshal(query)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return err
	}

	// Tạo yêu cầu xóa bằng truy vấn
	req := esapi.DeleteByQueryRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(string(doc)),
	}

	// Thực hiện xóa bằng truy vấn
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to delete all documents: %s", res.String())
	}

	return nil
}
