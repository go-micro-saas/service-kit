package main

import (
	"context"
	"flag"
	"fmt"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	"github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdlog "log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	serverNameSuffix = "-service"
)

var (
	configPath string
	sourceDir  string
	storeDir   string
)

func init() {
	flag.StringVar(&configPath, "consul_config", "", "consul config path, eg: -consul_config ./configs")
	flag.StringVar(&sourceDir, "source_dir", "", "store source path, eg: -source_dir path/to/source_dir")
	flag.StringVar(&storeDir, "store_dir", "", "custom store path, eg: -store_dir project_name/service_name/store_dir")
}

func currentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	}()

	// 配置
	confPath := configPath
	if confPath == "" {
		err = errorpkg.ErrorBadRequest("请配置consul config")
		panic(err)
	}
	if !filepath.IsAbs(confPath) {
		confPath = filepath.Join(currentPath(), confPath)
	}
	bootConfig, err := configutil.LoadingFile(confPath)
	if err != nil {
		panic(err)
	}

	// consul
	sourcePath := sourceDir
	if sourceDir == "" {
		err = errorpkg.ErrorBadRequest("请配置资源目录：source_dir")
		panic(err)
	}
	if !filepath.IsAbs(sourceDir) {
		sourcePath = filepath.Join(currentPath(), sourceDir)
	}
	consulHandler, err := NewConsulConfig(bootConfig, sourcePath)
	if err != nil {
		return
	}

	// 开始配置
	stdlog.Println("|==================== 更新配置到Consul 开始 ====================|")
	defer stdlog.Println("|==================== 更新配置到Consul 结束 ====================|")
	stdlog.Println("|*** Consul链接配置路径：	", confPath)
	stdlog.Println("|*** 资源配置路径：	", sourcePath)
	err = consulHandler.StoreToConsul()
	if err != nil {
		return
	}
}

// ConsulConfig ...
type ConsulConfig struct {
	cc         *api.Client
	sourcePath string
}

// NewConsulConfig 初始化
func NewConsulConfig(config *configpb.Bootstrap, sourcePath string) (*ConsulConfig, error) {
	if config.GetConsul() == nil {
		e := errorpkg.ErrorBadRequest("请先配置Consul配置再试")
		return nil, errorpkg.WithStack(e)
	}
	cc, err := consulpkg.NewClient(consulutil.ToConsulConfig(config.GetConsul()))
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return &ConsulConfig{
		cc:         cc,
		sourcePath: sourcePath,
	}, nil
}

// StoreToConsul 存到consul
func (s *ConsulConfig) StoreToConsul() error {
	configDataM, err := s.ReadConfigFiles()
	if err != nil {
		return err
	}
	ctx := context.Background()
	opt := &api.WriteOptions{}
	opt = opt.WithContext(ctx)
	for key := range configDataM {
		stdlog.Println("|*** Consul配置文件：", key)
		kv := &api.KVPair{
			Key:   key,
			Value: configDataM[key],
		}
		_, err := s.cc.KV().Put(kv, opt)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return errorpkg.WithStack(e)
		}
	}
	return nil
}

// ReadConfigFiles 读取文件
func (s *ConsulConfig) ReadConfigFiles() (map[string][]byte, error) {
	fs, err := filepathpkg.ReadDir(s.sourcePath)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	consulPath := storeDir
	if consulPath == "" {
		bs, err := configutil.LoadingFile(s.sourcePath)
		if err != nil {
			e := errorpkg.ErrorBadRequest(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		consulPath = apputil.Path(bs.GetApp())
	}
	if consulPath == "" {
		e := errorpkg.ErrorBadRequest("请配置存储路径：store_dir")
		return nil, errorpkg.WithStack(e)
	}
	stdlog.Println("|*** Consul存储路径：", consulPath)
	configDataM := make(map[string][]byte)
	for i := range fs {
		if fs[i].IsDir() {
			continue
		}
		destPath := filepath.Join(consulPath, fs[i].Name())
		filePath := filepath.Join(s.sourcePath, fs[i].Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		configDataM[destPath] = content
		//fmt.Println(destPath, len(content))
	}
	return configDataM, nil
}
