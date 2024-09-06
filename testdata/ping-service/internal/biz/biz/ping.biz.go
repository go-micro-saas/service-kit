package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	resourcev1 "github.com/go-micro-saas/service-kit/testdata/ping-service/api/ping/v1/resources"
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

func (s *pingBiz) TestingRequest(ctx context.Context) error {
	pingHTTPClient, err := clientutil.NewPingHTTPClient(s.serviceAPIManager, "ping-service-http")
	if err != nil {
		return err
	}
	pingGRPCClient, err := clientutil.NewPingGRPCClient(s.serviceAPIManager, "ping-service-grpc")
	if err != nil {
		return err
	}
	pingHTTPReq := &resourcev1.PingReq{Message: "request_by_http"}
	pingHTTPResp, err := pingHTTPClient.Ping(ctx, pingHTTPReq)
	if err != nil {
		return err
	}
	s.log.Infow("==> TestingRequest: ", pingHTTPResp.GetData().GetMessage())
	pingGRPCReq := &resourcev1.PingReq{Message: "request_by_grpc"}
	pingGRPCResp, err := pingGRPCClient.Ping(ctx, pingGRPCReq)
	if err != nil {
		return err
	}
	s.log.Infow("==> TestingRequest: ", pingGRPCResp.GetData().GetMessage())
	return nil
}
