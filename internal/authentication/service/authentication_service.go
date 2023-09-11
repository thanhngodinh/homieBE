package service

import (
	"context"
	"database/sql"
	"hostel-service/internal/authentication/domain"
	"hostel-service/internal/authentication/port"
)

type AuthenticationService interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (int64, error)
}

func NewAuthenticationService(db *sql.DB, repository port.AuthenticationRepository) AuthenticationService {
	return &authenticationService{db: db, repository: repository}
}

type authenticationService struct {
	db         *sql.DB
	repository port.AuthenticationRepository
}

func (s *authenticationService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.repository.GetByUsername(ctx, username)
}

func (s *authenticationService) Create(ctx context.Context, user *domain.User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Create(ctx, user)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
