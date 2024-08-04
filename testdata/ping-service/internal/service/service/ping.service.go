package pingservice

import (
	"context"
	bizrepo "github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/repo"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/service/dto"
	resourcev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	servicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

type pingService struct {
	servicev1.UnimplementedSrvPingServer

	pingBiz bizrepo.PingBizRepo
}

func NewPingService(pingBiz bizrepo.PingBizRepo) servicev1.SrvPingServer {
	return &pingService{
		pingBiz: pingBiz,
	}
}

func (s *pingService) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	param := dto.PingDTO.ToBoGetPingMessageParam(req)
	reply, err := s.pingBiz.GetPingMessage(ctx, param)
	if err != nil {
		return nil, err
	}
	resp := dto.PingDTO.ToPbPingResp(reply)
	return resp, err
}
