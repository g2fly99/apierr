package apierr

import (
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// // NewError 新构造一个Error
// func NewError(code, msg string) error {
// 	return errors.New(http.StatusOK, code, msg)
// }

type ErrorEncoder struct {
	whitelist []string
}

type ErrorEncoderOption func(e *ErrorEncoder)

func WithErrorEncoderWhitelist(pathPrefixList []string) ErrorEncoderOption {
	return func(e *ErrorEncoder) {
		e.whitelist = make([]string, 0)
		e.whitelist = append(e.whitelist, pathPrefixList...)
	}
}

// ErrorEncoder Kratos错误应答编码插件，将错误转换成CloudAPIv3格式
func ErrorEncoderFunc(l log.Logger, opts ...ErrorEncoderOption) khttp.EncodeErrorFunc {
	h := log.NewHelper(l)
	enc := &ErrorEncoder{
		whitelist: make([]string, 0),
	}
	for _, o := range opts {
		o(enc)
	}
	return func(w http.ResponseWriter, r *http.Request, err error) {
		h.Errorf("ErrorEncoder, %+v", err)
		if err == nil {
			return
		}

		// 请求在白名单中 , 跳过
		path := r.URL.Path
		for _, p := range enc.whitelist {
			if p == path {
				return
			} else if len(p) > 0 && p[len(p)-1] == '/' && strings.Contains(path, p) {
				return
			}
		}

		// 默认的错误应答
		reply := &APIErrorReply{
			Response: APIErrorResponse{
				Error: Error{
					Code:    ErrorInternalError,
					Message: err.Error(),
				},
			},
		}

		codec, _ := khttp.CodecForRequest(r, "Accept")
		// 拿到error并转换成kratos Error实体
		se := errors.FromError(err)
		if se != nil { // 正常来说，进入ErrorEncoder，
			// 通过Request Header的Accept中提取出对应的编码器

			w.Header().Set("Content-Type", contentType(codec.Name()))

			switch se.GetCode() {
			case http.StatusOK: // 业务错误码
				reply.Response.Error.Code = se.GetReason()
				reply.Response.Error.Message = se.GetMessage()
			case http.StatusUnauthorized:
				reply.Response.Error.Code = ErrorAuthFailure
				reply.Response.Error.Message = se.GetReason() + se.Message
			case http.StatusBadRequest:
				reply.Response.Error.Code = ErrorInvalidRequest
				reply.Response.Error.Message = se.GetReason() + se.Message
			default:
				reply.Response.Error.Code = http.StatusText(int(se.Code))
				reply.Response.Error.Message = se.GetReason() + se.Message
			}

			if reply.Response.Error.Message == "" {
				reply.Response.Error.Message = mapErrorCode[reply.Response.Error.Code]
			}

			bReply, _ := codec.Marshal(reply)
			w.WriteHeader(http.StatusOK)
			w.Write(bReply)
		} else { // 不符合规范的error
			reply := &APIErrorReply{
				Response: APIErrorResponse{
					Error: Error{
						Code:    ErrorInternalError,
						Message: err.Error(),
					},
				},
			}
			bReply, _ := codec.Marshal(reply)
			w.WriteHeader(http.StatusOK)
			w.Write(bReply)
		}

	}
}

// ContentType returns the content-type with base prefix.
func contentType(subtype string) string {
	baseContentType := "application"
	return strings.Join([]string{baseContentType, subtype}, "/")
}
