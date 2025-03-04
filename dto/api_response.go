package dto

type ApiResponse struct {
	Data         any       `json:"data,omitempty"`
	Status       int       `json:"status"`
	Message      string    `json:"message,omitempty"`
	ErrorType    ErrorType `json:"error_type,omitempty"`
	ErrorMessage string    `json:"error_message,omitempty"`
}

type ErrorType string

const (
	ErrorNone              ErrorType = ""
	ErrorExec              ErrorType = "execution"
	ErrorBadData           ErrorType = "bad_data"
	ErrorInternal          ErrorType = "internal"
	ErrorUnavailable       ErrorType = "unavailable"
	ErrorNotFound          ErrorType = "not_found"
	ErrorNotImplemented    ErrorType = "not_implemented"
	ErrorUnauthorized      ErrorType = "unauthorized"
	ErrorForbidden         ErrorType = "forbidden"
	ErrorAlreadyRegistered ErrorType = "already_registered"
)

type ApiError struct {
	Typ          ErrorType
	Err          error
	ErrorMessage string
}

func InternalError(err error) *ApiError {
	return &ApiError{
		Typ:          ErrorInternal,
		Err:          err,
		ErrorMessage: "internal server error",
	}
}
