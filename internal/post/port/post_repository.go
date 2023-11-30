package port

import (
	"context"
	"hostel-service/internal/post/domain"
)

type PostRepository interface {
	GetPosts(ctx context.Context, hostel *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, id string) (*domain.Post, error)
	GetPostByIds(ctx context.Context, ids []string) ([]domain.Post, error)
	CreatePost(ctx context.Context, hostel *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, hostel *domain.Post) (int64, error)
	DeletePost(ctx context.Context, hostel *domain.Post) (int64, error)
}
