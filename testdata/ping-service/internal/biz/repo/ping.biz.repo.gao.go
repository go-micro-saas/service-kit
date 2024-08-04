package bizrepo

import (
	"context"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/bo"
)

type PingBizRepo interface {
	GetPingMessage(ctx context.Context, param *bo.GetPingMessageParam) (*bo.GetPingMessageResult, error)
}
