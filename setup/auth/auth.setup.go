package authutil

import (
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/redis/go-redis/v9"
)

type authManager struct {
	conf                *configpb.Encrypt_TokenEncrypt
	redisCC             redis.UniversalClient
	loggerForMiddleware log.Logger

	// 不要直接使用 s.tokenXxx, 请使用 GetAuthorizationManager()
	tokenManager     authpkg.TokenManger
	tokenAuthRepo    authpkg.AuthRepo
	tokenManagerOnce sync.Once
}

type AuthorizationManager struct {
	TokenManager authpkg.TokenManger
	AuthManager  authpkg.AuthRepo
}

type AuthManager interface {
	GetTokenManger() (authpkg.TokenManger, error)
	GetAuthManger() (authpkg.AuthRepo, error)
}

func NewAuthManager(conf *configpb.Encrypt_TokenEncrypt, redisCC redis.UniversalClient, loggerForMiddleware log.Logger) (AuthManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt")
		return nil, errorpkg.WithStack(e)
	}
	if conf.GetSignKey() == "" {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt.sign_key")
		return nil, errorpkg.WithStack(e)
	}
	if conf.GetRefreshKey() == "" {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt.refresh_key")
		return nil, errorpkg.WithStack(e)
	}
	return &authManager{
		conf:                conf,
		redisCC:             redisCC,
		loggerForMiddleware: loggerForMiddleware,
	}, nil
}

func (s *authManager) GetAuthorizationManager() (*AuthorizationManager, error) {
	err := s.loadingTokenManagerOnce()
	if err != nil {
		return nil, err
	}
	return &AuthorizationManager{
		TokenManager: s.tokenManager,
		AuthManager:  s.tokenAuthRepo,
	}, nil
}

func (s *authManager) GetTokenManger() (authpkg.TokenManger, error) {
	manager, err := s.GetAuthorizationManager()
	if err != nil {
		return nil, err
	}
	return manager.TokenManager, nil
}
func (s *authManager) GetAuthManger() (authpkg.AuthRepo, error) {
	manager, err := s.GetAuthorizationManager()
	if err != nil {
		return nil, err
	}
	return manager.AuthManager, nil
}

func (s *authManager) loadingTokenManagerOnce() error {
	var err error
	s.tokenManagerOnce.Do(func() {
		s.tokenManager, s.tokenAuthRepo, err = s.loadingTokenManager()
		if err != nil {
			s.tokenManagerOnce = sync.Once{}
		}
	})
	return err
}

func (s *authManager) loadingTokenManager() (authpkg.TokenManger, authpkg.AuthRepo, error) {
	stdlog.Println("|*** LOADING: TokenManger: ...")
	tokenManger := authpkg.NewTokenManger(s.loggerForMiddleware, s.redisCC, authpkg.CheckAuthCacheKeyPrefix(nil))
	config := &authpkg.Config{
		SignCrypto:    authpkg.NewSignEncryptor(s.conf.GetSignKey()),
		RefreshCrypto: authpkg.NewCBCCipher(s.conf.GetRefreshKey()),
	}
	stdlog.Println("|*** LOADING: AuthManger: ...")
	authRepo, err := authpkg.NewAuthRepo(*config, s.loggerForMiddleware, tokenManger)
	if err != nil {
		return nil, nil, err
	}
	return tokenManger, authRepo, nil
}
