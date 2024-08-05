/*
//go:build wireinject
// +build wireinject
*/
package pingexport

import (
	serverutil "github.com/go-micro-saas/service-kit/server"
	"github.com/google/wire"
)

func a() (interface{}, error) {
	panic(wire.Build(
		serverutil.NewGRPCServer, serverutil.NewHTTPServer,
	))
	return nil, nil
}
