package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	"github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	serverNameSuffix = "-service"
)

var (
	configFlag string
	storePath  string
)

func init() {
	flag.StringVar(&configFlag, "conf", "", "config path, eg: -conf ./configs")
	flag.StringVar(&storePath, "path", "", "custom store path, eg: -path project_name/service_name/file_path")
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
	configPath := filepath.Join(currentPath(), "configs")
	bootConfig, err := configutil.LoadingFile(configPath)
	if err != nil {
		panic(err)
	}

	// consul
	absPath := configPath
	//absPath, err := filepath.Abs(configFlag)
	if err != nil {
		panic(err)
	}
	consulHandler, err := NewConsulConfig(bootConfig, absPath)
	if err != nil {
		return
	}

	// 开始配置
	stdlog.Println("|==================== 更新配置到Consul 开始 ====================|")
	defer stdlog.Println("|==================== 更新配置到Consul 结束 ====================|")
	err = consulHandler.StoreToConsul()
	if err != nil {
		return
	}
}

// loadingConfig 加载配置
func loadingConfig() (*configpb.Bootstrap, error) {
	handler := config.New(config.WithSource(file.NewSource(configFlag)))
	err := handler.Load()
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var conf = &configpb.Bootstrap{}
	err = handler.Scan(conf)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	// App配置
	if conf.App == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : app")
		return nil, err
	}

	// 服务配置
	if conf.Server == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : server")
		return nil, err
	}
	return conf, nil
}

// ConsulConfig ...
type ConsulConfig struct {
	config     *configpb.Bootstrap
	cc         *api.Client
	path       string
	serverName string
}

// NewConsulConfig 初始化
func NewConsulConfig(config *configpb.Bootstrap, absPath string) (*ConsulConfig, error) {
	if config.GetConsul() == nil {
		err := pkgerrors.New("请先配置Consul配置再试")
		return nil, err
	}
	cc, err := consulpkg.NewClient(consulutil.ToConsulConfig(config.GetConsul()))
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var serverName = config.GetApp().ServerName
	if serverName == "" {
		err = fmt.Errorf("查找不到服务名；配置路径: %s， 查找的服务名后缀：%s", configFlag, serverName)
		return nil, pkgerrors.WithStack(err)
	}

	return &ConsulConfig{
		config:     config,
		cc:         cc,
		path:       absPath,
		serverName: serverName,
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
			return pkgerrors.WithStack(err)
		}
	}
	return nil
}

// ReadConfigFiles 读取文件
func (s *ConsulConfig) ReadConfigFiles() (map[string][]byte, error) {
	fs, err := filepathpkg.ReadDir(s.path)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if s.serverName != s.config.App.ServerName {
		format := `配置中的服务名与配置路径中的服务名不一致；
	配置中的服务名：%s；
	配置路径：%s；
	配置路中的服务名：%s；`
		err = fmt.Errorf(format, s.config.App.ServerName, configFlag, s.serverName)
		return nil, pkgerrors.WithStack(err)
	}
	consulPath := storePath
	if consulPath == "" {
		consulPath = apputil.Path(s.config.GetApp())
	}
	stdlog.Println("|*** 本地配置路径：	", configFlag)
	stdlog.Println("|*** Consul配置路径：", consulPath)
	configDataM := make(map[string][]byte)
	for i := range fs {
		if fs[i].IsDir() {
			continue
		}
		destPath := filepath.Join(consulPath, fs[i].Name())
		filePath := filepath.Join(s.path, fs[i].Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, pkgerrors.WithStack(err)
		}
		configDataM[destPath] = content
		//fmt.Println(destPath, len(content))
	}
	return configDataM, nil
}
