package service

import (
	"context"
	"errors"
	"hostel-service/internal/admin/domain"
	"hostel-service/internal/admin/port"
	rate_port "hostel-service/internal/rate/port"
	"hostel-service/pkg/send_email"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetByUsername(ctx context.Context, username string) (*domain.Admin, error)
	GetAdminProfile(ctx context.Context, adminId string) (*domain.Admin, error)
	UpdatePassword(ctx context.Context, adminId string, oldPassword string, newPassword string) error

	SearchPosts(ctx context.Context, post *domain.PostFilter) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, postId string) (*domain.Post, error)
	UpdatePostStatus(ctx context.Context, id string, status string) (int64, error)

	SearchUsers(ctx context.Context, filter *domain.UserFilter) ([]domain.User, int64, error)
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	UpdateUserStatus(ctx context.Context, userId string, status string) error
	ResetPassword(ctx context.Context, userId string) error
}

func NewAdminService(
	adminRepo port.AdminRepository,
	rateRepo rate_port.RateRepository,
) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		rateRepo:  rateRepo,
	}
}

type adminService struct {
	adminRepo port.AdminRepository
	rateRepo  rate_port.RateRepository
}

func (s *adminService) GetByUsername(ctx context.Context, username string) (*domain.Admin, error) {
	return s.adminRepo.GetByUsername(ctx, username)
}

func (s *adminService) GetAdminProfile(ctx context.Context, adminId string) (*domain.Admin, error) {
	return s.adminRepo.GetById(ctx, adminId)
}

func (s *adminService) UpdatePassword(ctx context.Context, adminId string, oldPassword string, newPassword string) error {
	admin, err := s.adminRepo.GetById(ctx, adminId)
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldPassword)) != nil {
		return errors.New("password mismatch")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.adminRepo.UpdatePassword(ctx, adminId, string(hashedPassword))
}

func (s *adminService) SearchPosts(ctx context.Context, post *domain.PostFilter) ([]domain.Post, int64, error) {
	return s.adminRepo.GetPosts(ctx, post)
}

func (s *adminService) GetPostById(ctx context.Context, postId string) (*domain.Post, error) {
	res, err := s.adminRepo.GetPostById(ctx, postId)
	if err != nil {
		return nil, err
	}
	res.RateInfo, err = s.rateRepo.GetPostRate(ctx, postId)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *adminService) UpdatePostStatus(ctx context.Context, id string, status string) (int64, error) {
	return s.adminRepo.UpdatePostStatus(ctx, id, status)
}

// User
func (s *adminService) SearchUsers(ctx context.Context, filter *domain.UserFilter) ([]domain.User, int64, error) {
	return s.adminRepo.SearchUsers(ctx, filter)
}

func (s *adminService) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	return s.adminRepo.GetUserById(ctx, id)
}

func (s *adminService) ResetPassword(ctx context.Context, userId string) error {
	user, err := s.adminRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	newPassword := uuid.New().String()[0:6]

	err = send_email.SendPasswordResetEmail(user.Email, newPassword)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.adminRepo.UpdatePassword(ctx, userId, string(hashedPassword))
}

func (s *adminService) UpdateUserStatus(ctx context.Context, userId string, status string) error {
	return s.adminRepo.UpdateUserStatus(ctx, userId, status)
}
