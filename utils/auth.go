package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIdFromContext(ctx *gin.Context) uuid.UUID {
	userId := ctx.MustGet("user_id")

	return uuid.MustParse(userId.(string))
}
