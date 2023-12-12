package apierr

// ErrorCode 错误码结构
type ErrorCode struct {
	StatusCode int
	Code       string
	Message    string
}

var (
	ErrorHello = ErrorCode{10, "", ""}
)

const (
	ErrorInternalError = "InternalError" // 内部错误、未知错误

	// 登录态相关错误码， 未登录、登录态失效、登录态校验失败、登录态过期、签名错误

	// **** 身份认证与鉴权 ****
	ErrorAuthFailure         = "AuthFailure"          // 认证与鉴权，不能细分类型的错误。
	ErrorAuthFailureSession  = "AuthFailure.Session"  // 认证与鉴权，会话登录态异常。
	ErrorAuthFailureBusiness = "AuthFailure.Business" // 认证与鉴权，业务登录态异常。
	// ErrorAuthFailureUnauthorized = "AuthFailure.Unauthorized" // 用户未登录、登录态失效，或者接口为携带鉴权信息。
	// ErrorAuthFailureTokenFailure     = "AuthFailure.TokenFailure"     // 登录态校验失败，Cookie&Token校验失败
	// ErrorAuthFailureTokenExpire      = "AuthFailure.TokenExpire"      // 登录态过期
	// ErrorAuthFailureSignatureFailure = "AuthFailure.SignatureFailure" // 网关签名错误

	// **** 验证码 ****
	ErrorCaptcha                     = "CaptchaFailure"                      // 验证码错误
	ErrorCaptchaNotMatch             = "CaptchaFailure.Notmatch"             // 验证码错误, 验证码不匹配
	ErrorCaptchaRequestLimitExceeded = "CaptchaFailure.RequestLimitExceeded" // 验证码请求过于频繁

	// **** 客户端错误，请求错误****
	ErrorInvalidRequest              = "InvalidRequest"               // 请求错误
	ErrorInvalidRequestRepeated      = "InvalidRequest.Repeated"      // 请求被重放，触发防重放
	ErrorInvalidRequestLimitExceeded = "InvalidRequest.LimitExceeded" // 请求限频

	// **** 客户端错误，请求错误，参数错误 ****
	ErrorInvalidParameter             = "InvalidParameter"              // 缺少参数或参数不合法
	ErrorInvalidParameterMissingParam = "InvalidParameter.MissingParam" // 缺少参数
	ErrorInvalidParameterInvalidValue = "InvalidParameter.InvalidValue" // 参数值不合法

	// **** 服务端错误 ****
	ErrorOperation                 = "OperationFailure"                  // 操作失败, 原因未知
	ErrorOperationResourceNotFound = "OperationFailure.ResourceNotFound" // 资源不存在
	ErrorOperationUserNotFound     = "OperationFailure.UserNotFound"     // 用户不存在
	ErrorOperationUnauth           = "OperationFailure.UserNotFound"     // 用户不存在
	ErrorOperationUnauthorized     = "OperationFailure.Unauthorized"     // 未经授权的操作，操作越权
)

// mapErrorCode :
var mapErrorCode = map[string]string{
	ErrorInternalError: "内部错误，未知错误",

	// **** 身份认证与鉴权 ****
	ErrorAuthFailure:         "鉴权未通过",
	ErrorAuthFailureSession:  "鉴权未通过, 缺少会话态或会话态失效",
	ErrorAuthFailureBusiness: "鉴权未通过, 缺少登录态或登录态失效",
	// ErrorAuthFailureUnauthorized: "鉴权未通过",
	// ErrorAuthFailureTokenFailure:     "鉴权未通过，鉴权信息错误",
	// ErrorAuthFailureTokenExpire:      "鉴权未通过，鉴权信息过期",
	// ErrorAuthFailureSignatureFailure: "签名错误",

	// **** 验证码 ****
	ErrorCaptcha:                     "验证码错误",
	ErrorCaptchaNotMatch:             "验证码错误",
	ErrorCaptchaRequestLimitExceeded: "验证码请求过于频繁",

	// **** 客户端错误，请求错误****
	ErrorInvalidRequest:              "客户端错误",
	ErrorInvalidRequestRepeated:      "重放请求",
	ErrorInvalidRequestLimitExceeded: "请求超过频率",

	// **** 客户端错误，请求错误，参数错误 ****
	ErrorInvalidParameter:             "参数错误",
	ErrorInvalidParameterMissingParam: "参数错误，缺少参数",
	ErrorInvalidParameterInvalidValue: "参数错误，参数值不合法",

	// **** 服务端错误 ****
	ErrorOperation:                 "操作失败",
	ErrorOperationResourceNotFound: "操作失败，资源不存在",
	ErrorOperationUserNotFound:     "操作失败，用户不存在",
	ErrorOperationUnauthorized:     "操作失败，未经授权的操作",
}
