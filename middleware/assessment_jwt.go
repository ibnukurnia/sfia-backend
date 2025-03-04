package middleware

import (
	"fmt"
	"strings"
	"sv-sfia/dto"
	"sv-sfia/dto/responses"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func AssessmentJWT(checkParticipantFunc func(uuid.UUID) *dto.ApiError) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			zap.L().Error("missing authorization header")

			responses.ResponseError(ctx, &dto.ApiError{
				ErrorMessage: "missing authorization header",
				Typ:          dto.ErrorUnauthorized,
			})
			ctx.Abort()
			return
		}

		splits := strings.Split(bearerToken, " ")

		if len(splits) < 2 {
			responses.ResponseError(ctx, &dto.ApiError{
				Err:          fmt.Errorf("invalid token authorization"),
				ErrorMessage: "invalid token",
				Typ:          dto.ErrorUnauthorized,
			})
			ctx.Abort()
			return
		}

		tokenString := splits[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(""), nil
		})

		if err != nil {
			zap.L().Error("error parsing JWT", zap.Error(err))

			responses.ResponseError(ctx, &dto.ApiError{
				Err:          err,
				ErrorMessage: "invalid token",
				Typ:          dto.ErrorUnauthorized,
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(string)
			if !ok {
				zap.L().Error("missing user_id in claims")

				responses.ResponseError(ctx, &dto.ApiError{
					ErrorMessage: "missing user_id in token",
					Typ:          dto.ErrorUnauthorized,
				})
				ctx.Abort()
				return
			}

			userUuid, err := uuid.Parse(userID)
			if err != nil {
				zap.L().Error("missing user_id in claims")

				responses.ResponseError(ctx, &dto.ApiError{
					ErrorMessage: "failed parsing uuid",
					Typ:          dto.ErrorUnauthorized,
					Err:          err,
				})
				ctx.Abort()
				return
			}

			apiErr := checkParticipantFunc(userUuid)

			if apiErr != nil {
				responses.ResponseError(ctx, apiErr)
				ctx.Abort()

				return
			}

			ctx.Set("user_id", userID)
		} else {
			zap.L().Error("invalid token claims")

			responses.ResponseError(ctx, &dto.ApiError{
				ErrorMessage: "invalid token claims",
				Typ:          dto.ErrorUnauthorized,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
