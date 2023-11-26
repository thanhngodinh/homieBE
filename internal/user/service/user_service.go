package service

import (
	"context"
	"errors"
	send_email "hostel-service/internal/package/sendEmail"
	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/port"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UpdateUserSuggest(ctx context.Context, userUpdate *domain.UpdateUserSuggest) error
	GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error)
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
	GetByUsername(ctx context.Context, username string) (*domain.UserProfile, error)
	UpdatePassword(ctx context.Context, userId string, oldPassword string, newPassword string) error
	ResetPassword(ctx context.Context, userId string) error
	UpdateUserStatus(ctx context.Context, userId string, status string) error
	Create(ctx context.Context, user *domain.User) error
	SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error)
	GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error)
}

func NewUserService(
	userRepo port.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo port.UserRepository
}

func (s *userService) SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error) {
	return s.userRepo.SearchRoommates(ctx, filter)
}

func (s *userService) GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error) {
	return s.userRepo.GetUserProfile(ctx, userId)
}

func (s *userService) GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error) {
	return s.userRepo.GetRoommateById(ctx, userId)
}

func (s *userService) UpdateUserSuggest(ctx context.Context, userUpdate *domain.UpdateUserSuggest) error {
	return s.userRepo.UpdateUserSuggest(ctx, userUpdate)
}

func (s *userService) GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error) {
	return s.userRepo.GetUserSuggest(ctx, userId)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*domain.UserProfile, error) {
	return s.userRepo.GetByUsername(ctx, username)
}

func (s *userService) GetById(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetById(ctx, id)
}

func (s *userService) Create(ctx context.Context, user *domain.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) UpdatePassword(ctx context.Context, userId string, oldPassword string, newPassword string) error {
	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)) != nil {
		return errors.New("password mismatch")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.userRepo.UpdatePassword(ctx, userId, string(hashedPassword))
}

func (s *userService) ResetPassword(ctx context.Context, userId string) error {
	user, err := s.userRepo.GetById(ctx, userId)
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
	return s.userRepo.UpdatePassword(ctx, userId, string(hashedPassword))
}

func (s *userService) UpdateUserStatus(ctx context.Context, userId string, status string) error {
	return s.userRepo.UpdateUserStatus(ctx, userId, status)
}
