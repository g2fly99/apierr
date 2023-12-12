package apierr

// Error 实现了
type Error struct {
	Code    string
	Message string
}

type APIErrorResponse struct {
	Error Error
}

type APIErrorReply struct {
	Response APIErrorResponse
}
