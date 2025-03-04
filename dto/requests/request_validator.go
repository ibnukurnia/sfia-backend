package requests

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/responses"
	"sv-sfia/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type requestValidation struct {
	validator *validator.Validate
	ctx       *gin.Context
}

func NewValidation(ctx *gin.Context) *requestValidation {
	return &requestValidation{
		validator: validator.New(),
		ctx:       ctx,
	}
}

func (rV *requestValidation) Validate(request interfaces.RequestValidated) bool {
	err := rV.ctx.BindJSON(&request)

	if err != nil {
		zap.L().Error("error parse request json", zap.Error(err))

		responses.ResponseError(rV.ctx, &dto.ApiError{
			Typ: dto.ErrorInternal,
			Err: fmt.Errorf("failed to parse request json"),
		})

		return false
	}

	err = rV.validator.Struct(request)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		customMessages := request.Messages()

		for _, e := range validationErrors {
			field := e.StructField()
			tag := e.Tag()
			key := fmt.Sprintf("%s.%s", field, tag)

			if msg, exists := customMessages[key]; exists {
				responses.ResponseError(rV.ctx, &dto.ApiError{
					Typ:          dto.ErrorBadData,
					Err:          fmt.Errorf("%s", msg),
					ErrorMessage: msg,
				})

				return false
			}
		}
	}

	return true
}
