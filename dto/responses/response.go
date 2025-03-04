package responses

import (
	"net/http"
	"sv-sfia/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func WriteApiResponse(ctx *gin.Context, data any, message string, status int) {
	ctx.JSON(status, dto.ApiResponse{
		Data:    data,
		Status:  status,
		Message: message,
	})
}

func ResponseError(ctx *gin.Context, apiErr *dto.ApiError) {
	var code int
	switch apiErr.Typ {
	case dto.ErrorBadData, dto.ErrorAlreadyRegistered:
		code = http.StatusBadRequest
	case dto.ErrorExec:
		code = 422
	case dto.ErrorInternal:
		code = http.StatusInternalServerError
	case dto.ErrorNotFound:
		code = http.StatusNotFound
	case dto.ErrorNotImplemented:
		code = http.StatusNotImplemented
	case dto.ErrorUnauthorized:
		code = http.StatusUnauthorized
	case dto.ErrorForbidden:
		code = http.StatusForbidden
	default:
		code = http.StatusInternalServerError
	}

	zap.L().Error("returning response error", zap.Error(apiErr.Err))

	ctx.JSON(code, dto.ApiResponse{
		Status:       code,
		ErrorMessage: apiErr.ErrorMessage,
		ErrorType:    apiErr.Typ,
	})
}
