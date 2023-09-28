package app

import (
	"context"

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
)

type ApplicationContext struct {
	Hostel hostelPort.HostelHandler
	User   userPort.UserHandler
	My     myPort.MyHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	gormDb, err := gorm.Open(postgres.Open(conf.DB), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	logError := log.LogError
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	// Repo
	hostelRepository := hostelRepository.NewHostelAdapter(gormDb)
	userRepository := userRepository.NewUserAdapter(gormDb)
	myRepository := myRepository.NewMyAdapter(gormDb)

	hostelService := hostelService.NewHostelService(hostelRepository, userRepository)
	hostelHandler := hostelHandler.NewHostelHandler(hostelService, validator.Validate, logError)

	userService := userService.NewUserService(userRepository, hostelRepository)
	userHandler := userHandler.NewUserHandler(userService, validator.Validate, logError)

	myService := myService.NewMyService(myRepository, hostelRepository)
	myHandler := myHandler.NewMyHandler(myService, validator.Validate, logError)

	return &ApplicationContext{
		Hostel: hostelHandler,
		User:   userHandler,
		My:     myHandler,
	}, nil
}
