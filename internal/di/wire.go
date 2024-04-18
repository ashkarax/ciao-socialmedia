package di

import (
	"fmt"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	server "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/api"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/db"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	JWTmiddleware "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/middleware"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase"
	aws "github.com/ashkarax/ciao-socialmedia/pkg/aws_s3"
	gosmtp "github.com/ashkarax/ciao-socialmedia/pkg/go_smtp"
)

func InitializeAPI(config *config.Config) (*server.ServerHttp, error) {
	DB, err := db.ConnectDatabase(&config.DB)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM WIRE.GO IN DI")
		return nil, err
	}

	gosmtp.SmtpConfigsForEmailOtp(config.Smtp)
	aws.AWSS3FileUploaderSetup(config.AwsS3)

	jwtRepo := repository.NewJWTRepo(DB)
	jwtUseCase := usecase.NewJWTUseCase(jwtRepo)
	jwtMiddleWare := JWTmiddleware.NewJWTMiddleware(jwtUseCase, &config.Token)

	postRepository := repository.NewPostRepo(DB)
	postUseCase := usecase.NewPostUseCase(postRepository)
	postHandler := handler.NewPostHandler(postUseCase)

	relationRepository := repository.NewRelationRepo(DB)
	relationUseCase := usecase.NewRelationUseCase(relationRepository)
	relationHandler := handler.NewRelationHandler(relationUseCase)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &config.Token, postRepository, relationRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	auth2oHandler := handler.NewAuth2oHandler(&config.Auth)

	serverHttp := server.NewServerHttp(&config.ApiKey, &config.PortMngr, jwtMiddleWare, userHandler, postHandler, relationHandler,auth2oHandler)
	return serverHttp, nil
}
