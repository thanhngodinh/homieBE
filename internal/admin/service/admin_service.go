package service

import (
	"context"
	"errors"
	"hostel-service/internal/admin/domain"
	"hostel-service/internal/admin/port"
	post_port "hostel-service/internal/post/port"
	user_port "hostel-service/internal/user/port"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetByUsername(ctx context.Context, username string) (*domain.Admin, error)
	UpdatePassword(ctx context.Context, adminId string, oldPassword string, newPassword string) error
}

func NewAdminService(
	adminRepo port.AdminRepository,
	postRepo post_port.PostRepository,
	userRepo user_port.UserRepository,

) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		postRepo:  postRepo,
		userRepo:  userRepo,
	}
}

type adminService struct {
	adminRepo port.AdminRepository
	postRepo  post_port.PostRepository
	userRepo  user_port.UserRepository
}

func (s *adminService) GetByUsername(ctx context.Context, username string) (*domain.Admin, error) {
	return s.adminRepo.GetByUsername(ctx, username)
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
