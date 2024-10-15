package events

import (
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/go-micro-saas/service-kit/testdata/test-service/internal/biz/repo"
)

type testingEvent struct {
	log *log.Helper
}

func NewTestingEvent(
	logger log.Logger,
) bizrepos.TestingEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/event"))

	return &testingEvent{
		log: logHelper,
	}
}
