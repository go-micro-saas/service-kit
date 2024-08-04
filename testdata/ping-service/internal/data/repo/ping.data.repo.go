package datarepo

import (
	"context"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/po"
)

type PingDataRepo interface {
	GetMockPingMessage(ctx context.Context, param *po.MockPingMessageParam) (*po.MockPingMessageReply, error)
}
