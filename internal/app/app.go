package app

import (
	"context"

	sv "github.com/core-go/core"
	v "github.com/core-go/core/v10"
	"github.com/core-go/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	hostelHandler "hostel-service/internal/hostel/adapter/handler"
	hostelRepository "hostel-service/internal/hostel/adapter/repository"
	hostelPort "hostel-service/internal/hostel/port"
	hostelService "hostel-service/internal/hostel/service"

	userHandler "hostel-service/internal/user/adapter/handler"
	userRepository "hostel-service/internal/user/adapter/repository"
	userPort "hostel-service/internal/user/port"
	userService "hostel-service/internal/user/service"

	myHandler "hostel-service/internal/my/adapter/handler"
	myRepository "hostel-service/internal/my/adapter/repository"
	myPort "hostel-service/internal/my/port"
	myService "hostel-service/internal/my/service"

	authHandler "hostel-service/internal/authentication/adapter/handler"
	authRepository "hostel-service/internal/authentication/adapter/repository"
	authPort "hostel-service/internal/authentication/port"
	authService "hostel-service/internal/authentication/service"
)

type ApplicationContext struct {
	Hostel hostelPort.HostelHandler
	User   userPort.UserHandler
	My     myPort.MyHandler
	Auth   authPort.AuthenticationHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	gormDb, err := gorm.Open(postgres.Open(conf.DB), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	logError := log.LogError
	status := sv.InitializeStatus(conf.Status)
	action := sv.InitializeAction(conf.Action)
	validator := v.NewValidator()

	hostelRepository := hostelRepository.NewHostelAdapter(gormDb)
	hostelService := hostelService.NewHostelService(hostelRepository)
	hostelHandler := hostelHandler.NewHostelHandler(hostelService, validator.Validate, logError)

	userRepository := userRepository.NewUserAdapter(gormDb)
	userService := userService.NewUserService(userRepository, hostelRepository)
	userHandler := userHandler.NewUserHandler(userService, validator.Validate, logError)

	myRepository := myRepository.NewMyAdapter(gormDb)
	myService := myService.NewMyService(myRepository, hostelRepository)
	myHandler := myHandler.NewMyHandler(myService, validator.Validate, logError)

	authRepository := authRepository.NewAuthenticationAdapter(gormDb)
	authService := authService.NewAuthenticationService(gormDb, authRepository)
	authHandler := authHandler.NewAuthenticationHandler(authService, status, logError, validator.Validate, &action)

	return &ApplicationContext{
		Hostel: hostelHandler,
		Auth:   authHandler,
		User:   userHandler,
		My:     myHandler,
	}, nil
}
