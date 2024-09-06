package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/bo"
	bizrepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/repo"
	datarepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/repo"
)

type pingBiz struct {
	log               *log.Helper
	serviceAPIManager clientutil.ServiceAPIManager
	pingData          datarepo.PingDataRepo
}

func NewPingBiz(
	logger log.Logger,
	serviceAPIManager clientutil.ServiceAPIManager,
	pingData datarepo.PingDataRepo,
) bizrepo.PingBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "ping/biz/ping"))

	return &pingBiz{
		log:      logHelper,
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
