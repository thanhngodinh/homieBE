package service

import (
	"context"
	"errors"
	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/port"
	"hostel-service/pkg/send_email"
	"hostel-service/pkg/send_otp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*domain.UserProfile, string, int, error)
	Register(ctx context.Context, username string, phone string, name string) error
	SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error)
	GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error)
	GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error)
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
	UpdateUserSuggest(ctx context.Context, userUpdate *domain.UpdateUserSuggest) error
	UpdatePassword(ctx context.Context, userId string, oldPassword string, newPassword string) error
	VerifyPhone(ctx context.Context, userId string, phone string) error
	VerifyPhoneOTP(ctx context.Context, userId string, otp string) error
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

func (s *userService) Login(ctx context.Context, username string, password string) (*domain.UserProfile, string, int, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, "", 0, errors.New("Not found user. Error: " + err.Error())
	}
	if user.Status != "A" {
		return nil, "", 0, errors.New("User not active. Error: " + err.Error())
	}

	if password == "" || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, "", 0, errors.New("password not match")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  user.Id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString(domain.USER_SECRET_KEY)
	if err != nil {
		return nil, "", -1, err
	}

	if user.IsVerifiedEmail == nil || !*user.IsVerifiedEmail {
		return user, token, 2, nil
	}

	return user, token, 1, nil

}

func (s *userService) Register(ctx context.Context, username string, phone string, name string) error {
	now := time.Now()
	f := false
	userId := "user-" + uuid.NewString()
	user := &domain.CreateUser{
		Id:              userId,
		Username:        username,
		Email:           username,
		Phone:           phone,
		Name:            name,
		IsVerifiedEmail: &f,
		IsVerifiedPhone: &f,
		CreatedAt:       &now,
		CreatedBy:       userId,
	}

	password := uuid.New().String()[0:8]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = send_email.SendVerifyEmail(username, password)
	if err != nil {
		return err
	}

	return s.userRepo.Create(ctx, user)
}

func (s *userService) UpdatePassword(ctx context.Context, userId string, oldPassword string, newPassword string) error {
	user, err := s.userRepo.GetUserById(ctx, userId)
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

func (s *userService) VerifyPhone(ctx context.Context, userId string, phone string) error {
	expirationTime := time.Now().Add(2 * time.Minute)
	otp := send_otp.GenerateOTP()

	err := send_otp.SendOTP("+84"+phone, otp)
	if err != nil {
		return err
	}

	return s.userRepo.VerifyPhone(ctx, userId, phone, otp, expirationTime)
}

func (s *userService) VerifyPhoneOTP(ctx context.Context, userId string, otp string) error {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	if user.ExpirationTime == nil || user.OTP == "" {
		return errors.New("Can't process verify")
	}

	if otp != user.OTP || time.Now().After(*user.ExpirationTime) {
		return errors.New("OTP is invalid")
	}

	return s.userRepo.VerifyPhoneOTP(ctx, userId, otp)
}
