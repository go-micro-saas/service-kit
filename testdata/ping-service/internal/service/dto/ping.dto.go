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

func (p *pingDTO) ToPbPingRespData(dataModel *bo.GetPingMessageResult) *resourcev1.PingRespData {
	res := &resourcev1.PingRespData{
		Message: dataModel.Message,
	}
	return res
}
