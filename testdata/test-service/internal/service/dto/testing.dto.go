package dto

import (
	"github.com/go-micro-saas/service-kit/testdata/test-service/internal/biz/bo"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	TestingDto testing
)

type testing struct{}

func (s *testing) ToBoXxx() *bo.Testdata {
	res := &bo.Testdata{}
	return res
}

func (s *testing) ToPbXxx() *emptypb.Empty {
	res := &emptypb.Empty{}

	return res
}
