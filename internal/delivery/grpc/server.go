package grpc

import (
	"errors"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	"github.com/prometheus/common/log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"github.com/valyala/fasthttp/reuseport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// OmdbClientService interface to hold all the GRPC methods
type OmdbClientService struct {
	cfg               *config.Config
	localTime         *time.Location
	omdbClientUseCase usecase.OmdbClientUseCase
}

//GRPCServer is a gRPC server struct
type GRPCServer struct {
	server  *grpc.Server
	address string
}

//Option is struct
type Option struct {
	ListenAddress string
}

//New is method
func New(o *Option, conf *config.Config, usecase usecase.UseCase) (grpcServer *GRPCServer, err error) {
	if conf == nil {
		err = errors.New("[grpc][New] invalid config given")
	}
	if o == nil {
		err = errors.New("[grpc][New] invalid grpc option given")
	}

	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return grpcServer, err
	}

	omdbClientService := OmdbClientService{
		cfg:               conf,
		omdbClientUseCase: usecase.OmdbClientUseCase,
		localTime:         localTime,
	}

	server := grpc.NewServer()

	pb.RegisterOmdbClientServer(server, omdbClientService)
	reflection.Register(server)

	grpcServer = &GRPCServer{
		server:  server,
		address: o.ListenAddress,
	}
	log.Infoln("[Init][GRPC] new grpc success")
	return grpcServer, err
}

//Start is method
func (gs *GRPCServer) Start() error {
	var err error
	gs.signalInterrupt()
	go func() {
		l, err := reuseport.Listen("tcp4", gs.address)
		if err != nil {
			l, err = net.Listen("tcp", gs.address)
			if err != nil {
				return
			}
		}

		log.Infof("[Init][GRPC] starting grpc server on %+v\n", gs.address)
		err = gs.server.Serve(l)
		if err != nil {
			return
		}
	}()
	return err
}

//signalInterrupt is method
func (gs *GRPCServer) signalInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	// Block until a signal is received.
	go func() {
		for {
			select {
			case _, ok := <-c:
				if !ok {
					return
				}
				gs.server.GracefulStop()
				log.Infoln("GRPC server stopped")
			}
		}
	}()
}
