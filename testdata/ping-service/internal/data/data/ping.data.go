package data

import (
	"context"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/po"
	datarepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/repo"
)

type pingData struct{}

func NewPingData() datarepo.PingDataRepo {
	return &pingData{}
}

func (p *pingData) GetMockPingMessage(ctx context.Context, param *po.MockPingMessageParam) (*po.MockPingMessageReply, error) {
	return &po.MockPingMessageReply{Message: "mock request message: " + param.Message}, nil
}
