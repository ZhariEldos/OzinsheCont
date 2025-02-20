package Structs

type ApiError struct {
	Error string
	Where string
}

func NewApiError(msg string, where string) *ApiError {
	return &ApiError{msg, where}
}
