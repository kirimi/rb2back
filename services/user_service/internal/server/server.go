package server

import (
	"context"
	"fmt"
	"github.com/kirimi/rb2backend/user_service/internal/api"
	"github.com/kirimi/rb2backend/user_service/internal/config"
	pb "github.com/kirimi/rb2backend/user_service/pkg/generated/user_service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer func(l net.Listener) {
		_ = l.Close()
	}(l)

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: cfg.Grpc.GetMaxConnectionIdle(),
			Timeout:           cfg.Grpc.GetTimeout(),
			MaxConnectionAge:  cfg.Grpc.GetMaxConnectionAge(),
			Time:              cfg.Grpc.GetTimeout(),
		}),
		//grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		//	grpc_ctxtags.UnaryServerInterceptor(),
		//	grpc_prometheus.UnaryServerInterceptor,
		//	grpc_opentracing.UnaryServerInterceptor(),
		//	grpcrecovery.UnaryServerInterceptor(),
		//)),
	)

	pb.RegisterUserServiceServer(grpcServer, api.NewUserAPI())

	go func() {
		log.Info().Msgf("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
	log.Info().Msgf("grpcServer shut down correctly")

	return nil
}
