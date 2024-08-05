package main

import (
	"flag"
	setuputil "github.com/go-micro-saas/service-kit/setup"
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
	launcher, err := setuputil.NewLauncherManager(flagconf)
	if err != nil {
		stdlog.Fatalf("setuputil.NewLauncherManager faild, error: %+v", err)
	}
	_ = launcher
}
