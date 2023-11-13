package app

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postHandler "hostel-service/internal/post/adapter/handler"
	postRepository "hostel-service/internal/post/adapter/repository"
	postPort "hostel-service/internal/post/port"
	postService "hostel-service/internal/post/service"

	userHandler "hostel-service/internal/user/adapter/handler"
	userRepository "hostel-service/internal/user/adapter/repository"
	userPort "hostel-service/internal/user/port"
	userService "hostel-service/internal/user/service"

	myHandler "hostel-service/internal/my/adapter/handler"
	myRepository "hostel-service/internal/my/adapter/repository"
	myPort "hostel-service/internal/my/port"
	myService "hostel-service/internal/my/service"

	utilitiesHandler "hostel-service/internal/utilities/adapter/handler"
	utilitiesRepository "hostel-service/internal/utilities/adapter/repository"
	utilitiesPort "hostel-service/internal/utilities/port"
	utilitiesService "hostel-service/internal/utilities/service"

	rateHandler "hostel-service/internal/rate/adapter/handler"
	rateRepository "hostel-service/internal/rate/adapter/repository"
	ratePort "hostel-service/internal/rate/port"
	rateService "hostel-service/internal/rate/service"

	chatHandler "hostel-service/internal/chat/adapter/handler"
	// chatRepository "hostel-service/internal/chat/adapter/repository"
	chatPort "hostel-service/internal/chat/port"
	chatService "hostel-service/internal/chat/service"
)

type ApplicationContext struct {
	Post      postPort.PostHandler
	Utilities utilitiesPort.UtilitiesHandler
	User      userPort.UserHandler
	My        myPort.MyHandler
	Rate      ratePort.RateHandler
	Chat      chatPort.ChatHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	gormDb, err := gorm.Open(postgres.Open(conf.DB), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	esCfg := elasticsearch.Config{
		CloudID: "dbf646249c124369bf2e94bafc31712a:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJGUzYWFlYzIyZTAwNDRmMjNiZjEzMmEzYjEzZjE4ZmM5JDVlY2VjZGExOTIwZjRjMmQ4MzU1MTUwNTcxMDQzOTRl",
		APIKey:  "UmdxOWtJc0IzVjdNaEowcDJNQ2w6ak8tMGNYSHlRczZGZlRJZlNhYlZ6QQ==",
	}
	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	validate := validator.New()

	// Repo
	postRepository := postRepository.NewPostAdapter(gormDb)
	userRepository := userRepository.NewUserRepo(gormDb)
	myRepository := myRepository.NewMyAdapter(gormDb)
	utilitiesRepository := utilitiesRepository.NewUtilitiesAdapter(gormDb)
	rateRepository := rateRepository.NewRateAdapter(gormDb)

	postService := postService.NewPostService(postRepository, userRepository, rateRepository, es)
	postHandler := postHandler.NewPostHandler(postService, validate)

	utilitiesService := utilitiesService.NewUtilitiesService(utilitiesRepository)
	utilitiesHandler := utilitiesHandler.NewUtilitiesHandler(utilitiesService, validate)

	userService := userService.NewUserService(userRepository, postRepository)
	userHandler := userHandler.NewUserHandler(userService, validate)

	myService := myService.NewMyService(myRepository, postRepository)
	myHandler := myHandler.NewMyHandler(myService, validate)

	rateService := rateService.NewRateService(rateRepository)
	rateHandler := rateHandler.NewRateHandler(rateService, validate)

	chatService := chatService.NewChatService()
	chatHandler := chatHandler.NewChatHandler(chatService)

	return &ApplicationContext{
		Post:      postHandler,
		Utilities: utilitiesHandler,
		User:      userHandler,
		My:        myHandler,
		Rate:      rateHandler,
		Chat:      chatHandler,
	}, nil
}
