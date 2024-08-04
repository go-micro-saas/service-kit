package biz

import (
	"context"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/bo"
	bizrepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/repo"
	datarepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/repo"
)

type pingBiz struct {
	pingData datarepo.PingDataRepo
}

func NewPingBiz(pingData datarepo.PingDataRepo) bizrepo.PingBizRepo {
	return &pingBiz{
		pingData: pingData,
	}
}

func (s *pingBiz) GetPingMessage(ctx context.Context, param *bo.GetPingMessageParam) (*bo.GetPingMessageResult, error) {
	dataParam := param.ToPoMockPingMessageParam()
	reply, err := s.pingData.GetMockPingMessage(ctx, dataParam)
	if err != nil {
		return nil, err
	}
	res := &bo.GetPingMessageResult{}
	res.SetByPoMockPingMessageReply(reply)
	return res, nil
}
