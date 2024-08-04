package dto

import (
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/bo"
	resourcev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
)

var (
	PingDTO pingDTO
)

type pingDTO struct{}

func (p *pingDTO) ToBoGetPingMessageParam(req *resourcev1.PingReq) *bo.GetPingMessageParam {
	res := &bo.GetPingMessageParam{
		Message: req.Message,
	}
	return res
}

func (p *pingDTO) ToPbPingResp(dataModel *bo.GetPingMessageResult) *resourcev1.PingResp {
	res := &resourcev1.PingResp{
		Message: dataModel.Message,
	}
	return res
}
