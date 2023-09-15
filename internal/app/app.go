package app

import (
	"context"

	sv "github.com/core-go/core"
	v "github.com/core-go/core/v10"
	"github.com/core-go/health"
	"github.com/core-go/log"
	q "github.com/core-go/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	hostelHandler "hostel-service/internal/hostel/adapter/handler"
	hostelRepository "hostel-service/internal/hostel/adapter/repository"
	hostelPort "hostel-service/internal/hostel/port"
	hostelService "hostel-service/internal/hostel/service"

	authHandler "hostel-service/internal/authentication/adapter/handler"
	authRepository "hostel-service/internal/authentication/adapter/repository"
	authPort "hostel-service/internal/authentication/port"
	authService "hostel-service/internal/authentication/service"
)

type ApplicationContext struct {
	Health *health.Handler
	Hostel hostelPort.HostelHandler
	Auth   authPort.AuthenticationHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}

	gormDb, err := gorm.Open(postgres.Open(conf.DB), &gorm.Config{})

	logError := log.LogError
	status := sv.InitializeStatus(conf.Status)
	action := sv.InitializeAction(conf.Action)
	validator := v.NewValidator()

	hostelRepository := hostelRepository.NewHostelAdapter(db)
	hostelService := hostelService.NewHostelService(db, hostelRepository)
	hostelHandler := hostelHandler.NewHostelHandler(hostelService, validator.Validate, logError)
	if err != nil {
		return nil, err
	}

	authRepository := authRepository.NewAuthenticationAdapter(gormDb)
	authService := authService.NewAuthenticationService(gormDb, authRepository)
	authHandler := authHandler.NewAuthenticationHandler(authService, status, logError, validator.Validate, &action)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health: healthHandler,
		Hostel: hostelHandler,
		Auth:   authHandler,
	}, nil
}
