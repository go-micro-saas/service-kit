package service

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	resourcev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
)

type HomeService struct{}

func NewRootPath() *HomeService {
	return &HomeService{}
}

func (s *HomeService) Homepage(w http.ResponseWriter, r *http.Request) {
	data := &resourcev1.PingResp{
		Data: &resourcev1.PingRespData{
			Message: "Hello World!",
		},
	}
	err := apppkg.SuccessResponseEncoder(w, r, data)
	if err != nil {
		apppkg.ErrorResponseEncoder(w, r, err)
	}
}
