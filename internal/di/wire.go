package di

import (
	"fmt"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	server "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/api"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/db"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase"
	gosmtp "github.com/ashkarax/ciao-socialmedia/pkg/go_smtp"
)

func InitializeAPI(config *config.Config) (*server.ServerHttp, error) {
	DB, err := db.ConnectDatabase(&config.DB)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM WIRE.GO IN DI")
		return nil, err
	}

	gosmtp.SmtpConfigsForEmailOtp(config.Smtp)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &config.Token)
	userHandler := handler.NewUserHandler(userUseCase)

	serverHttp := server.NewServerHttp(&config.PortMngr, userHandler)
	return serverHttp, nil
}
