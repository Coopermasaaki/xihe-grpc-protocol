package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/grpc/competition"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/evaluate"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/inference"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/training"
	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

func NewTrainingClient(endpoint string) (*TrainingClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &TrainingClient{
		clientConn: &c,
		cli:        protocol.NewTrainingClient(c.conn),
	}, nil
}

func NewInferenceClient(endpoint string) (*InferenceClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &InferenceClient{
		clientConn: &c,
		cli:        protocol.NewInferenceClient(c.conn),
	}, nil
}

func NewEvaluateClient(endpoint string) (*EvaluateClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &EvaluateClient{
		clientConn: &c,
		cli:        protocol.NewEvaluateClient(c.conn),
	}, nil
}

func NewCompetitionClient(endpoint string) (*CompetitionClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &CompetitionClient{
		clientConn: &c,
		cli:        protocol.NewCompetitionClient(c.conn),
	}, nil
}

func newConn(endpoint string) (c clientConn, err error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err == nil {
		c.conn = conn
	}

	return
}

type clientConn struct {
	conn *grpc.ClientConn
}

func (c *clientConn) Disconnect() error {
	if c == nil || c.conn == nil {
		return nil
	}

	err := c.conn.Close()
	c.conn = nil

	return err
}

type TrainingClient struct {
	*clientConn

	cli protocol.TrainingClient
}

func (c *TrainingClient) SetTrainingInfo(index *training.TrainingIndex, info *training.TrainingInfo) error {
	_, err := c.cli.SetTrainingInfo(
		context.Background(),
		&protocol.TrainingInfo{
			Id:            index.Id,
			User:          index.User,
			Status:        info.Status,
			LogPath:       info.LogPath,
			Duration:      int32(info.Duration),
			ProjectId:     index.ProjectId,
			AimZipPath:    info.AimZipPath,
			OutputZipPath: info.OutputZipPath,
		},
	)

	return err
}

type InferenceClient struct {
	*clientConn

	cli protocol.InferenceClient
}

func (c *InferenceClient) SetInferenceInfo(index *inference.InferenceIndex, info *inference.InferenceInfo) error {
	_, err := c.cli.SetInferenceInfo(
		context.Background(),
		&protocol.InferenceInfo{
			Id:         index.Id,
			User:       index.User,
			ProjectId:  index.ProjectId,
			LastCommit: index.LastCommit,
			Error:      info.Error,
			AccessUrl:  info.AccessURL,
		},
	)

	return err
}

type EvaluateClient struct {
	*clientConn

	cli protocol.EvaluateClient
}

func (c *EvaluateClient) SetEvaluateInfo(index *evaluate.EvaluateIndex, info *evaluate.EvaluateInfo) error {
	_, err := c.cli.SetEvaluateInfo(
		context.Background(),
		&protocol.EvaluateInfo{
			Id:         index.Id,
			User:       index.User,
			ProjectId:  index.ProjectId,
			TrainingId: index.TrainingID,
			Error:      info.Error,
			AccessUrl:  info.AccessURL,
		},
	)

	return err
}

type CompetitionClient struct {
	*clientConn

	cli protocol.CompetitionClient
}

func (c *CompetitionClient) SetSubmissionInfo(
	index *competition.SubmissionIndex, info *competition.SubmissionInfo,
) error {
	_, err := c.cli.SetSubmissionInfo(
		context.Background(),
		&protocol.SubmissionInfo{
			Id:            index.Id,
			Phase:         index.Phase,
			CompetitionId: index.CompetitionId,
			Score:         info.Score,
		},
	)

	return err
}
