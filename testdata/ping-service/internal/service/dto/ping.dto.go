package dto

import (
	resourcev1 "github.com/go-micro-saas/service-kit/testdata/ping-service/api/ping/v1/resources"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/bo"
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

func (p *pingDTO) ToPbPingRespData(dataModel *bo.GetPingMessageResult) *resourcev1.PingRespData {
	res := &resourcev1.PingRespData{
		Message: dataModel.Message,
	}
	return res
}
