package utils

import (
	"sv-sfia/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func ParseUUid(s string) (uuid.UUID, *dto.ApiError) {
	id, err := uuid.Parse(s)

	if err != nil {
		zap.L().Error("error parsing uuid", zap.Error(err))

		return id, &dto.ApiError{
			Typ:          dto.ErrorBadData,
			ErrorMessage: "Invalid uuid format",
			Err:          err,
		}
	}

	return id, nil
}
