package main

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository"
	"log"

	"github.com/gorilla/mux"

	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		cfg *config.Config
		err error
	)

	// initialize app config
	cfg, err = config.Init()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// initialize repository
	repo, err := repository.Init(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize repository %+v", err)
	}

	// initialize usecase
	useCase := usecase.Init(cfg, repo)

	//initialize gRPC
	initGRPC(useCase)

	// initialize delivery
	router := mux.NewRouter()
	api := delivery.Init(cfg, router, useCase)
	api.InitRoutes()

}

// initGRPC is used for initializing gRPC server
// placed here instead of in config.go to avoid import cycle
func initGRPC(usecase usecase.UseCase) {
	grpcServer, err := grpc.New(&grpc.Option{
		ListenAddress: ":" + config.CFG.Server.GRPCPort,
	}, config.CFG, usecase)
	if err != nil {
		log.Fatalf("[Init][GRPC] failed to Init GRPC %v\n", err)
	}
	err = grpcServer.Start()
	if err != nil {
		log.Fatalf("[Init][GRPC] failed to Start GRPC %v\n", err)
	}
}
