package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
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

type Exercise struct {
	ID        string `yaml:"id"`
	Lang      string `yaml:"lang"`
	Homedir   string `yaml:"homedir"`
	TestCases []TC   `yaml:"tests"`
}

type TC struct {
	Name     string `yaml:"name"`
	Input    string `yaml:"input"`
	Expected string `yaml:"expected"`
}

func (g *GraderClient) Snd(ctx context.Context, exer exercise.Exercise) (err error) {
	var (
		uf    = make([]*grader.UserFile, len(exer.Files))
		tests []byte
		file  *os.File
	)

	tests, err = yaml.Marshal(exer.Files)
	if err != nil {
		return fmt.Errorf("error marshalling yaml: %w", err)
	}

	for i, v := range exer.Files {
		var buf []byte

		file, err = os.Open(v.Path)
		if err != nil {
			log.Error().Err(err).Msg("Failed to open file")

			return fmt.Errorf("error opening file: %w", err)
		}

		buf, err = io.ReadAll(file)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read file")

			return
		}

		uf[i] = &grader.UserFile{
			Name:    v.Name,
			Content: buf,
		}
	}

	req := &grader.ExerciseRequest{
		Lang:  exer.Lang,
		Tests: tests,
		Files: uf,
	}

	resp, err := g.client.Exercise(ctx, req)
	if err != nil {
		return fmt.Errorf("grader error: %w", err)
	}

	log.Debug().Msgf("Received response: %v", resp)

	return nil
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
