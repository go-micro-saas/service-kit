package main

import (
	"flag"
	runservices "github.com/go-micro-saas/service-kit/testdata/ping-service/cmd/ping-service/run"
	stdlog "log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	app, cleanup, err := runservices.GetServerApp(flagconf)
	if err != nil {
		stdlog.Fatalf("==> runservices.GetServerApp failed: %+v\n", err)
	}
	defer func() {
		if cleanup != nil {
			cleanup()
		}
	}()
	// start
	if err := app.Run(); err != nil {
		stdlog.Fatalf("==> app.Run failed: %+v\n", err)
	}
}
