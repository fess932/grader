package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grader/pkg/exercise"
	"grader/pkg/gen/grader"
	"io"
	"os"
)

const chunkSize = 1024 * 64 // 64KB optimal chunk size for streaming binary data

func NewGraderClient(addr string) (*GraderClient, error) {
	log.Debug().Msg("Creating new grpc client")

	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grader connect error: %w", err)
	}

	return &GraderClient{
		client: grader.NewGraderServiceClient(grpcConn),
	}, nil
}

type GraderClient struct {
	client grader.GraderServiceClient
}

func (g GraderClient) SendExercise(exercise *exercise.Exercise) error {
	var (
		filePath = exercise.Files[0].Path
		file     *os.File

		buf = make([]byte, chunkSize)
		num int
	)

	stream, err := g.client.Upload(context.Background())
	if err != nil {
		return fmt.Errorf("grader upload error: %w", err)
	}

	file, err = os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}

	for {
		num, err = file.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("read file error: %w", err)
		}

		if err = stream.Send(&grader.UploadRequest{
			Chunk: buf[:num],
		}); err != nil {
			return fmt.Errorf("send file error: %w", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("close and recv error: %w", err)
	}

	log.Debug().Msgf("Upload result: %s", res.Name)

	return nil

}
