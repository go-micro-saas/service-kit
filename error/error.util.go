package errorutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdhttp "net/http"
	"strconv"
)

var (
	OK       int32 = 0
	StatusOK int32 = stdhttp.StatusOK
)

type Response interface {
	GetCode() int32
	GetReason() string
	GetMessage() string
}

type GetMetadata interface {
	GetMetadata() map[string]string
}

func CheckHTTPStatus(statusCode int) *errors.Error {
	if statusCode >= stdhttp.StatusOK && statusCode < stdhttp.StatusMultipleChoices {
		return nil
	}
	reason := ""
	if v, ok := errorpkg.ERROR_name[int32(statusCode)]; ok {
		reason = v
	} else {
		reason = "HTTP_CODE_" + strconv.Itoa(statusCode)
	}
	e := errors.New(statusCode, reason, stdhttp.StatusText(statusCode))
	return e
}

func CheckResponseCode(resp Response) *errors.Error {
	if resp.GetCode() == OK {
		return nil
	}
	e := errors.New(int(resp.GetCode()), resp.GetReason(), resp.GetMessage())
	if md, ok := resp.(GetMetadata); ok {
		e.Metadata = md.GetMetadata()
	}
	return e
}

func CheckResponseStatus(resp Response) *errors.Error {
	if resp.GetCode() == OK || resp.GetCode() == StatusOK {
		return nil
	}
	e := errors.New(int(resp.GetCode()), resp.GetReason(), resp.GetMessage())
	if md, ok := resp.(GetMetadata); ok {
		e.Metadata = md.GetMetadata()
	}
	return e
}
