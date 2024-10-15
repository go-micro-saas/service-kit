package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/go-micro-saas/service-kit/testdata/test-service/internal/biz/repo"
)

type testingBiz struct {
	log *log.Helper
}

func NewTestingBiz(
	logger log.Logger,
) bizrepos.TestingBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &testingBiz{
		log: logHelper,
	}
}
