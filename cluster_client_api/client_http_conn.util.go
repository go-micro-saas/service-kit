package clientutil

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func NewHTTPClient(otherOpts ...http.ClientOption) (*http.Client, error) {
	return nil, errorpkg.ErrorUnimplemented("")
}
