package response

import "errors"

type BaseResponse[T any] struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type MetaResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInvalidID       = errors.New("Bad Request")
	ErrProjectNotFound = errors.New("Project Not Found")
	InternalError      = errors.New("Internal Server Error")
)

type EmptyObj struct{}

func SuccessResponse[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{
		Status:  true,
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

func ErrorResponse(code int, msg string) MetaResponse {
	return MetaResponse{
		Status:  false,
		Code:    code,
		Message: msg,
	}
}
