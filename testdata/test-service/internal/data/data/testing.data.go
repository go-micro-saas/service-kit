package data

import (
	"github.com/go-kratos/kratos/v2/log"
	datarepos "github.com/go-micro-saas/service-kit/testdata/test-service/internal/data/repo"
)

type testingData struct {
	log *log.Helper
}

func NewTestingData(
	logger log.Logger,
) datarepos.TestingDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/data/data"))

	return &testingData{
		log: logHelper,
	}
}
