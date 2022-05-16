package delivery

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"grader/configs"
	"grader/pkg/gen/grader"
	"net"
)

type IGrader interface {
	Grade()
}

func NewGraderService(gu IGrader, config configs.GraderConfig) *GraderService {
	return &GraderService{gu: gu, addr: config.Addr}
}

type GraderService struct {
	grader.UnimplementedGraderServiceServer

	addr string
	gu   IGrader
}

func (s *GraderService) Upload(server grader.GraderService_UploadServer) error {
	log.Info().Msg("GraderService.Upload")
	for {
		req, err := server.Recv()
		if err != nil {
			log.Error().Err(err).Msg("GraderService.Upload.Recv")
			return err
		}
		log.Info().Msgf("GraderService.Upload.Recv: %v", req)
		//s.gu.Grade()
	}
}

func (s *GraderService) Run() error {
	server := grpc.NewServer()
	grader.RegisterGraderServiceServer(server, s)

	log.Debug().Msgf("Starting GraderService on %s", s.addr)

	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err = server.Serve(l); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	server.GracefulStop()

	return nil
}
