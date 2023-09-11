package service

import (
	"context"
	"database/sql"
	"log"
	"hostel-service/internal/authentication/adapter/repository"
	"hostel-service/internal/authentication/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestLoadByUserName(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewAuthenticationAdapter(db)
	defer repo.DB.Close()
	service := NewAuthenticationService(db, repo)
	password, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	now := time.Now()
	mUser := &domain.User{
		Id:        uuid.New().String(),
		Username:  "user1",
		Password:  string(password),
		CreatedAt: &now,
	}
	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).
		AddRow(mUser.Id, mUser.Username, mUser.Password, mUser.CreatedAt, mUser.UpdatedAt)
	mock.
		ExpectPrepare(`select * from users where username = ? limit 1`).
		ExpectQuery().
		WithArgs(mUser.Username).
		WillReturnRows(rows)
	user, err := service.GetByUsername(context.Background(), mUser.Username)
	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, mUser, user)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewAuthenticationAdapter(db)
	defer repo.DB.Close()
	service := NewAuthenticationService(db, repo)
	password, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	now := time.Now()
	mUser := &domain.User{
		Id:        uuid.New().String(),
		Username:  "user1",
		Password:  string(password),
		CreatedAt: &now,
	}
	mock.ExpectBegin()
	mock.ExpectPrepare(`insert into users values(?,?,?,?,?)`).ExpectExec().
		WithArgs(mUser.Id, mUser.Username, mUser.Password, mUser.CreatedAt, mUser.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	rowsAffected, err := service.Create(context.Background(), mUser)
	assert.Equal(t, int64(1), rowsAffected)
	assert.NoError(t, err)
}
